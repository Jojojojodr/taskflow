<script setup lang="ts">
import { ref } from 'vue'
import { logout } from '../utils/api'

// Mobile menu state
const isMobileMenuOpen = ref(false)

// Toggle mobile menu
const toggleMobileMenu = () => {
    isMobileMenuOpen.value = !isMobileMenuOpen.value
}

// Handle logout
const handleLogout = async () => {
    try {
        const result = await logout()
        if (result.success) {
            console.log('Logout successful')
            // Reload the page to trigger authentication check
            window.location.reload()
        } else {
            console.error('Logout failed:', result.error)
        }
    } catch (error) {
        console.error('Logout error:', error)
        // Even if logout API fails, reload to clear local state
        window.location.reload()
    }
}
</script>

<template>
    <nav class="bg-white shadow-lg border-b border-gray-200">
        <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
            <div class="flex justify-between items-center h-16">
                <!-- Logo/Brand -->
                <div class="flex-shrink-0">
                    <router-link to="/"
                        class="text-2xl font-bold text-gray-800 hover:text-blue-600 transition-colors duration-200">
                        TaskFlow
                    </router-link>
                </div>

                <!-- Desktop Navigation -->
                <div class="hidden md:flex md:items-center md:space-x-4">
                    <ul class="flex space-x-8">
                        <li>
                            <router-link to="/"
                                class="text-gray-600 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 hover:bg-blue-50 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                                Home
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/tasks"
                                class="text-gray-600 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 hover:bg-blue-50 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                                Tasks
                            </router-link>
                        </li>
                        <li>
                            <router-link to="/board"
                                class="text-gray-600 hover:text-blue-600 px-3 py-2 rounded-md text-sm font-medium transition-colors duration-200 hover:bg-blue-50 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                                Board
                            </router-link>
                        </li>
                    </ul>

                    <!-- Logout Button -->
                    <div class="ml-6 pl-6 border-l border-gray-200">
                        <button @click="handleLogout"
                            class="bg-red-600 hover:bg-red-700 text-white px-4 py-2 rounded-md text-sm font-medium transition-colors duration-200 flex items-center space-x-2">
                            <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1">
                                </path>
                            </svg>
                            <span>Logout</span>
                        </button>
                    </div>
                </div>

                <!-- Mobile menu button -->
                <div class="md:hidden">
                    <button @click="toggleMobileMenu"
                        class="text-gray-600 hover:text-blue-600 focus:outline-none focus:ring-2 focus:ring-blue-500 focus:ring-offset-2 p-2 rounded-md">
                        <svg class="h-6 w-6" :class="{ 'hidden': isMobileMenuOpen, 'block': !isMobileMenuOpen }"
                            fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M4 6h16M4 12h16M4 18h16" />
                        </svg>
                        <svg class="h-6 w-6" :class="{ 'block': isMobileMenuOpen, 'hidden': !isMobileMenuOpen }"
                            fill="none" viewBox="0 0 24 24" stroke="currentColor">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M6 18L18 6M6 6l12 12" />
                        </svg>
                    </button>
                </div>
            </div>
        </div>

        <!-- Mobile Navigation Menu -->
        <div class="md:hidden" :class="{ 'block': isMobileMenuOpen, 'hidden': !isMobileMenuOpen }">
            <div class="px-2 pt-2 pb-3 space-y-1 bg-gray-50 border-t border-gray-200">
                <router-link to="/" @click="toggleMobileMenu"
                    class="text-gray-600 hover:text-blue-600 hover:bg-blue-50 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                    Home
                </router-link>
                <router-link to="/tasks" @click="toggleMobileMenu"
                    class="text-gray-600 hover:text-blue-600 hover:bg-blue-50 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                    Tasks
                </router-link>
                <router-link to="/board" @click="toggleMobileMenu"
                    class="text-gray-600 hover:text-blue-600 hover:bg-blue-50 block px-3 py-2 rounded-md text-base font-medium transition-colors duration-200 router-link-active:text-blue-600 router-link-active:bg-blue-50">
                    Board
                </router-link>

                <!-- Mobile Logout Button -->
                <div class="border-t border-gray-200 pt-3 mt-3">
                    <button @click="handleLogout"
                        class="w-full bg-red-600 hover:bg-red-700 text-white px-3 py-2 rounded-md text-base font-medium transition-colors duration-200 flex items-center justify-center space-x-2">
                        <svg class="w-4 h-4" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                            <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                d="M17 16l4-4m0 0l-4-4m4 4H7m6 4v1a3 3 0 01-3 3H6a3 3 0 01-3-3V7a3 3 0 013-3h4a3 3 0 013 3v1">
                            </path>
                        </svg>
                        <span>Logout</span>
                    </button>
                </div>
            </div>
        </div>
    </nav>
</template>

<style scoped></style>