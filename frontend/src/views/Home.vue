<template>
    <div class="min-h-screen flex flex-col items-center justify-center px-4 py-8">
        <div class="max-w-4xl w-full text-center">
            <h1 class="text-4xl font-bold text-gray-800 mb-8">Welcome to TaskFlow</h1>

            <!-- API Health Status -->
            <div class="mb-8">
                <div class="flex items-center justify-center gap-2">
                    <span class="text-sm font-medium text-gray-600">API Status:</span>
                    <div v-if="healthLoading" class="flex items-center gap-2">
                        <div class="animate-spin rounded-full h-4 w-4 border-2 border-blue-500 border-t-transparent">
                        </div>
                        <span class="text-sm text-gray-500">Checking...</span>
                    </div>
                    <div v-else-if="apiHealthy === true" class="flex items-center gap-2">
                        <div class="w-3 h-3 bg-green-500 rounded-full"></div>
                        <span class="text-sm text-green-600 font-medium">Healthy</span>
                    </div>
                    <div v-else-if="apiHealthy === false" class="flex items-center gap-2">
                        <div class="w-3 h-3 bg-red-500 rounded-full"></div>
                        <span class="text-sm text-red-600 font-medium">Unavailable</span>
                    </div>
                </div>
            </div>

            <div class="flex flex-col md:flex-row gap-8 justify-center">
                <div class="bg-white shadow-lg rounded-lg p-6 flex-1 max-w-md">
                    <h2 class="text-2xl font-semibold text-gray-800 mb-4">Get Started</h2>
                    <p class="text-gray-600 mb-4">
                        Welcome to your task management application. Organize your work and boost your productivity.
                    </p>
                    <router-link to="/tasks"
                        class="inline-block bg-blue-500 hover:bg-blue-600 text-white font-bold py-2 px-4 rounded transition-colors">
                        View Tasks
                    </router-link>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from 'vue'
import { checkHealth } from '@/utils/api'

// Reactive state for API health
const apiHealthy = ref<boolean | null>(null)
const healthLoading = ref(false)
const healthData = ref<any>(null)

// Function to check API health
const performHealthCheck = async () => {
    healthLoading.value = true
    try {
        const result = await checkHealth()

        if (result.success) {
            apiHealthy.value = true
            healthData.value = result.data
            console.log('API Health Check Success:', result.data)
        } else {
            apiHealthy.value = false
            console.error('API Health Check Failed:', result.error)
        }
    } catch (error) {
        apiHealthy.value = false
        console.error('Health check error:', error)
    } finally {
        healthLoading.value = false
    }
}

// Call API on component mount
onMounted(() => {
    performHealthCheck()
})
</script>