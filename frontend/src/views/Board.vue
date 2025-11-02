<template>
    <div class="container mx-auto px-4 py-8">
        <div class="max-w-7xl mx-auto">
            <h1 class="text-3xl font-bold text-gray-800 mb-8">Task Board</h1>

            <div class="flex flex-col md:flex-row gap-6">
                <!-- To Do Column -->
                <div class="bg-gray-50 rounded-lg p-4 flex-1">
                    <h2 class="font-semibold text-gray-800 mb-4">To Do</h2>
                    <div class="space-y-3">
                        <div class="bg-white p-3 rounded shadow">
                            <h3 class="font-medium text-gray-800">Sample Task 1</h3>
                            <p class="text-sm text-gray-600 mt-1">Task description goes here</p>
                        </div>
                        <div class="bg-white p-3 rounded shadow">
                            <h3 class="font-medium text-gray-800">Sample Task 2</h3>
                            <p class="text-sm text-gray-600 mt-1">Another task description</p>
                        </div>
                    </div>
                </div>

                <!-- In Progress Column -->
                <div class="bg-blue-50 rounded-lg p-4 flex-1">
                    <h2 class="font-semibold text-gray-800 mb-4">In Progress</h2>
                    <div class="space-y-3">
                        <div class="bg-white p-3 rounded shadow">
                            <h3 class="font-medium text-gray-800">Working Task</h3>
                            <p class="text-sm text-gray-600 mt-1">Currently being worked on</p>
                        </div>
                    </div>
                </div>

                <!-- Done Column -->
                <div class="bg-green-50 rounded-lg p-4 flex-1">
                    <h2 class="font-semibold text-gray-800 mb-4">Done</h2>
                    <div class="space-y-3">
                        <div class="bg-white p-3 rounded shadow">
                            <h3 class="font-medium text-gray-800">Completed Task</h3>
                            <p class="text-sm text-gray-600 mt-1">This task is finished</p>
                        </div>
                    </div>
                </div>
            </div>
        </div>
    </div>
</template>

<script setup lang="ts">
// Board page component
</script>