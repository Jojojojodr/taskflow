<template>
    <div class="min-h-screen bg-gray-50 py-8">
        <div class="max-w-4xl mx-auto px-4 sm:px-6 lg:px-8">
            <!-- Header -->
            <div class="flex justify-between items-center mb-8">
                <h1 class="text-3xl font-bold text-gray-900">Tasks</h1>
                <button
                    class="bg-blue-600 hover:bg-blue-700 text-white px-4 py-2 rounded-lg font-medium transition-colors">
                    Add Task
                </button>
            </div>

            <!-- Search bar -->
            <div class="mb-6">
                <div class="relative">
                    <input type="text" v-model="searchQuery" placeholder="Search tasks..."
                        class="w-full pl-10 pr-4 py-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 focus:border-transparent" />
                    <div class="absolute inset-y-0 left-0 pl-3 flex items-center pointer-events-none">
                        <svg class="h-5 w-5 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z"></path>
                        </svg>
                    </div>
                </div>
            </div>

            <!-- Loading state -->
            <div v-if="loading" class="flex justify-center items-center py-12">
                <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600"></div>
            </div>

            <!-- Error state -->
            <div v-else-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4 mb-6">
                <div class="flex">
                    <div class="flex-shrink-0">
                        <svg class="h-5 w-5 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.728-.833-2.498 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z">
                            </path>
                        </svg>
                    </div>
                    <div class="ml-3">
                        <h3 class="text-sm font-medium text-red-800">Error loading tasks</h3>
                        <div class="mt-2 text-sm text-red-700">
                            {{ error }}
                        </div>
                        <div class="mt-4">
                            <button @click="loadTasks"
                                class="bg-red-600 hover:bg-red-700 text-white px-3 py-1 rounded text-sm">
                                Try Again
                            </button>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Tasks list -->
            <div v-else-if="filteredTasks.length > 0" class="space-y-4">
                <div v-for="task in filteredTasks" :key="task.id"
                    class="bg-white rounded-lg shadow-sm border border-gray-200 p-6 hover:shadow-md transition-shadow">
                    <div class="flex items-start justify-between">
                        <div class="flex items-start space-x-3">
                            <input type="checkbox" :checked="task.completed" @change="toggleTask(task)"
                                class="mt-1 h-4 w-4 text-blue-600 focus:ring-blue-500 border-gray-300 rounded" />
                            <div class="flex-1">
                                <h3 :class="[
                                    'text-lg font-medium',
                                    task.completed ? 'text-gray-500 line-through' : 'text-gray-900'
                                ]">
                                    {{ task.title }}
                                </h3>
                                <p v-if="task.description" :class="[
                                    'mt-1 text-sm',
                                    task.completed ? 'text-gray-400' : 'text-gray-600'
                                ]">
                                    {{ task.description }}
                                </p>
                            </div>
                        </div>
                    </div>
                </div>
            </div>

            <!-- Empty state -->
            <div v-else class="text-center py-12">
                <svg class="mx-auto h-12 w-12 text-gray-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                    <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M9 5H7a2 2 0 00-2 2v10a2 2 0 002 2h8a2 2 0 002-2V7a2 2 0 00-2-2h-2M9 5a2 2 0 002 2h2a2 2 0 002-2M9 5a2 2 0 012-2h2a2 2 0 012 2">
                    </path>
                </svg>
                <h3 class="mt-2 text-sm font-medium text-gray-900">No tasks found</h3>
                <p class="mt-1 text-sm text-gray-500">Get started by creating your first task.</p>
            </div>

            <!-- Task count -->
            <div v-if="tasks.length > 0" class="mt-8 text-center text-sm text-gray-500">
                Showing {{ filteredTasks.length }} of {{ tasks.length }} tasks
                <span v-if="completedCount > 0">({{ completedCount }} completed)</span>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, computed, onMounted } from 'vue'
import { getTasks, updateTask, type Task } from '../utils/api'

// Reactive state
const tasks = ref<Task[]>([])
const loading = ref(false)
const error = ref<string | null>(null)
const searchQuery = ref('')

// Computed properties
const filteredTasks = computed(() => {
    if (!searchQuery.value) return tasks.value
    const query = searchQuery.value.toLowerCase()
    return tasks.value.filter(task =>
        task.title.toLowerCase().includes(query) ||
        task.description.toLowerCase().includes(query)
    )
})

const completedCount = computed(() => {
    return tasks.value.filter(task => task.completed).length
})

// Methods
const loadTasks = async () => {
    loading.value = true
    error.value = null

    try {
        const response = await getTasks()

        if (response.success && response.data) {
            tasks.value = response.data.tasks || []
        } else {
            error.value = response.error || 'Failed to load tasks'
        }
    } catch (err) {
        console.error('Error loading tasks:', err)
        error.value = 'An unexpected error occurred'
    } finally {
        loading.value = false
    }
}

const toggleTask = async (task: Task) => {
    // Store the original state in case we need to revert
    const originalCompleted = task.completed

    // Optimistically update the UI
    task.completed = !task.completed

    try {
        const response = await updateTask(task.id, {
            completed: task.completed
        })

        if (!response.success) {
            // Revert the change if the API call failed
            task.completed = originalCompleted
            console.error('Failed to update task:', response.error)
            // You could show a toast notification here
        } else {
            console.log('Task updated successfully:', task.id, 'completed:', task.completed)
        }
    } catch (err) {
        // Revert the change if there was an error
        task.completed = originalCompleted
        console.error('Error updating task:', err)
    }
}

// Load tasks when component mounts
onMounted(() => {
    loadTasks()
})
</script>