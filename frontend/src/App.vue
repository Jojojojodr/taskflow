<template>
    <!-- Loading state -->
    <div v-if="isLoading"
    class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex items-center justify-center">
        <div class="text-center">
            <div class="animate-spin rounded-full h-12 w-12 border-b-2 border-blue-600 mx-auto mb-4"></div>
            <p class="text-gray-600">Checking authentication...</p>
        </div>
    </div>

    <!-- Show login if not authenticated -->
    <div v-else-if="!isAuthenticated">
        <main
        class="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100 flex flex-col items-center justify-center px-4 sm:px-6 lg:px-8">
            <!-- Welcome Section -->
            <div class="text-center mb-8">
                <div class="mb-6">
                    <svg class="mx-auto h-16 w-16 text-blue-600" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                        <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                        d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z">
                    </path>
                </svg>
                </div>
                <h1 class="text-4xl font-bold text-gray-900 mb-4">
                    Welcome to TaskFlow
                </h1>
                <p class="text-xl text-gray-600 mb-2">
                    You are not logged in
                </p>
                <p class="text-lg text-gray-500">
                    Please login or register to continue
                </p>
            </div>

            <!-- Login Component Container -->
            <div class="w-full max-w-md">
                <Login :onLoginSuccess="handleLoginSuccess" />
            </div>

            <!-- Additional Info -->
            <div class="mt-8 text-center">
                <p class="text-sm text-gray-500">
                    Organize your tasks efficiently with our powerful task management system
                </p>
            </div>
        </main>
    </div>

    <!-- Show main app if user is authenticated -->
    <div v-else>
        <header>
            <Navbar />
        </header>

        <main>
            <router-view />
        </main>
    </div>
</template>

<script setup lang="ts">
import { onMounted, ref } from 'vue'
import Navbar from './components/NavBarComponent.vue'
import Login from './components/AuthComponent.vue'
import { validateUser } from './utils/api'

const isAuthenticated = ref<boolean>(false)
const isLoading = ref<boolean>(true)

const checkAuthentication = async () => {
    isLoading.value = true
    console.log('🔍 Starting authentication check...')
    try {
        const result = await validateUser()
        console.log('🔍 Validation result:', result)
        if (result.success && result.data && result.data.user) {
            isAuthenticated.value = true
            console.log('✅ User is authenticated:', result.data)
        } else {
            isAuthenticated.value = false
            console.log('❌ User is not authenticated - no user data:', result)
        }
    } catch (error) {
        console.error('🚨 Authentication check failed:', error)
        isAuthenticated.value = false
    } finally {
        isLoading.value = false
        console.log('🔍 Final state - isAuthenticated:', isAuthenticated.value, 'isLoading:', isLoading.value)
    }
}

const handleLoginSuccess = (user: any) => {
    console.log('Login successful:', user)
    isAuthenticated.value = true
}

onMounted(() => {
    checkAuthentication()
})
</script>