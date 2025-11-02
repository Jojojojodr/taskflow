<template>
    <div class="flex items-center justify-center bg-gray-50 py-12 px-4 sm:px-6 lg:px-8">
        <div class="max-w-md w-full space-y-8">
            <div>
                <h2 class="mt-6 text-center text-3xl font-extrabold text-gray-900">
                    {{ isRegisterMode ? 'Create your account' : 'Sign in to your account' }}
                </h2>
                <p class="mt-2 text-center text-sm text-gray-600">
                    {{ isRegisterMode ? 'Already have an account?' : "Don't have an account?" }}
                    <button @click="toggleMode" class="font-medium text-blue-600 hover:text-blue-500 transition-colors">
                        {{ isRegisterMode ? 'Sign in' : 'Sign up' }}
                    </button>
                </p>
            </div>

            <form class="mt-8 space-y-6" @submit.prevent="handleSubmit">
                <div class="rounded-md shadow-sm -space-y-px">
                    <div>
                        <label for="name" class="sr-only">Name</label>
                        <input id="name" name="name" type="text" required v-model="formData.name"
                            class="relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 rounded-t-md focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                            placeholder="Name" />
                    </div>

                    <div v-if="isRegisterMode">
                        <label for="email" class="sr-only">Email</label>
                        <input id="email" name="email" type="email" v-model="formData.email"
                            class="relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm"
                            placeholder="Email (optional)" />
                    </div>

                    <div>
                        <label for="password" class="sr-only">Password</label>
                        <input id="password" name="password" type="password" required v-model="formData.password"
                            :class="[
                                'relative block w-full px-3 py-2 border border-gray-300 placeholder-gray-500 text-gray-900 focus:outline-none focus:ring-blue-500 focus:border-blue-500 focus:z-10 sm:text-sm',
                                isRegisterMode ? 'rounded-b-md' : 'rounded-b-md'
                            ]" placeholder="Password" />
                    </div>
                </div>

                <!-- Error message -->
                <div v-if="error" class="bg-red-50 border border-red-200 rounded-lg p-4">
                    <div class="flex">
                        <div class="flex-shrink-0">
                            <svg class="h-5 w-5 text-red-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M12 9v2m0 4h.01m-6.938 4h13.856c1.54 0 2.502-1.667 1.732-2.5L13.732 4c-.77-.833-1.728-.833-2.498 0L3.732 16.5c-.77.833.192 2.5 1.732 2.5z">
                                </path>
                            </svg>
                        </div>
                        <div class="ml-3">
                            <h3 class="text-sm font-medium text-red-800">
                                {{ isRegisterMode ? 'Registration failed' : 'Login failed' }}
                            </h3>
                            <div class="mt-2 text-sm text-red-700">
                                {{ error }}
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Success message -->
                <div v-if="successMessage" class="bg-green-50 border border-green-200 rounded-lg p-4">
                    <div class="flex">
                        <div class="flex-shrink-0">
                            <svg class="h-5 w-5 text-green-400" fill="none" stroke="currentColor" viewBox="0 0 24 24">
                                <path stroke-linecap="round" stroke-linejoin="round" stroke-width="2"
                                    d="M5 13l4 4L19 7"></path>
                            </svg>
                        </div>
                        <div class="ml-3">
                            <h3 class="text-sm font-medium text-green-800">Success!</h3>
                            <div class="mt-2 text-sm text-green-700">
                                {{ successMessage }}
                            </div>
                        </div>
                    </div>
                </div>

                <div>
                    <button type="submit" :disabled="loading"
                        class="group relative w-full flex justify-center py-2 px-4 border border-transparent text-sm font-medium rounded-md text-white bg-blue-600 hover:bg-blue-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-blue-500 disabled:opacity-50 disabled:cursor-not-allowed transition-colors">
                        <span v-if="loading" class="absolute left-0 inset-y-0 flex items-center pl-3">
                            <div class="animate-spin rounded-full h-4 w-4 border-b-2 border-white"></div>
                        </span>
                        {{ loading ? 'Please wait...' : (isRegisterMode ? 'Sign up' : 'Sign in') }}
                    </button>
                </div>
            </form>
        </div>
    </div>
</template>

<script setup lang="ts">
import { ref, reactive } from 'vue'
import { login, register, type LoginRequest, type RegisterRequest } from '../utils/api'

// Component props
interface Props {
    onLoginSuccess?: (user: any) => void
}

const props = withDefaults(defineProps<Props>(), {
    onLoginSuccess: () => { }
})

// Component state
const isRegisterMode = ref(false)
const loading = ref(false)
const error = ref<string | null>(null)
const successMessage = ref<string | null>(null)

const formData = reactive({
    name: '',
    email: '',
    password: ''
})

// Methods
const toggleMode = () => {
    isRegisterMode.value = !isRegisterMode.value
    error.value = null
    successMessage.value = null
    // Clear form data when switching modes
    formData.name = ''
    formData.email = ''
    formData.password = ''
}

const handleSubmit = async () => {
    error.value = null
    successMessage.value = null
    loading.value = true

    try {
        if (isRegisterMode.value) {
            const registerData: RegisterRequest = {
                name: formData.name,
                password: formData.password,
                email: formData.email || undefined
            }

            const result = await register(registerData)

            if (result.success && result.data) {
                // Attempt to auto-login immediately after registration
                const loginData: LoginRequest = {
                    name: formData.name,
                    password: formData.password
                }

                const loginResult = await login(loginData)

                if (loginResult.success && loginResult.data) {
                    successMessage.value = 'Registered and logged in!'
                    props.onLoginSuccess(loginResult.data.user)
                    // Clear form
                    formData.name = ''
                    formData.email = ''
                    formData.password = ''
                } else {
                    // If auto-login fails, fall back to login mode
                    successMessage.value = 'Account created successfully! Please sign in.'
                    isRegisterMode.value = false
                    formData.password = ''
                    try {
                        formData.name = result.data.user?.name || formData.name
                    } catch (e) {
                    }
                    formData.email = ''
                }
            } else {
                error.value = result.error || 'Registration failed'
            }
        } else {
            const loginData: LoginRequest = {
                name: formData.name,
                password: formData.password
            }

            const result = await login(loginData)

            if (result.success && result.data) {
                successMessage.value = 'Login successful!'
                props.onLoginSuccess(result.data.user)
                // Clear form
                formData.name = ''
                formData.password = ''
            } else {
                error.value = result.error || 'Login failed'
            }
        }
    } catch (err) {
        console.error('Authentication error:', err)
        error.value = 'An unexpected error occurred'
    } finally {
        loading.value = false
    }
}
</script>