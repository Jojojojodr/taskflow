import axios from "axios";

// TypeScript interfaces for API responses
export interface Task {
    id: string
    title: string
    description: string
    completed: boolean
}

export interface TasksResponse {
    tasks: Task[]
    count: number
}

export interface User {
    id: string
    name: string
    email?: string
}

export interface LoginRequest {
    name: string
    password: string
}

export interface RegisterRequest {
    name: string
    password: string
    email?: string
}

export interface AuthResponse {
    message: string
    user: User
}

const API_SERVER = import.meta.env.VITE_API_SERVER || import.meta.env.API_SERVER
const API_PORT = import.meta.env.VITE_API_PORT || import.meta.env.API_PORT

const API_BASE_URL = ''

console.log('API Configuration:', {
    DEV: import.meta.env.DEV,
    API_SERVER,
    API_PORT,
    API_BASE_URL
})

// Create axios instance with base configuration
const apiClient = axios.create({
    baseURL: API_BASE_URL,
    timeout: 5000,
    headers: {
        'Content-Type': 'application/json',
    },
    withCredentials: true,
})

// Health check function
export const checkHealth = async () => {
    try {
        const response = await apiClient.get('/api/v1/health')
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error) {
        console.error('Health check failed:', error)
        return {
            success: false,
            error: error instanceof Error ? error.message : 'Unknown error',
            status: axios.isAxiosError(error) ? error.response?.status : null
        }
    }
}

// Get all tasks function
// Get tasks from the API
export const getTasks = async () => {
    try {
        const response = await apiClient.get<TasksResponse>("/api/v1/tasks")
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (getTasks):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.message || error.message || "Failed to fetch tasks",
            status: error.response?.status || 500
        }
    }
}

// Update a task
export const updateTask = async (taskId: string, taskData: Partial<Task>) => {
    try {
        const response = await apiClient.put(`/api/v1/task?id=${taskId}`, taskData)
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (updateTask):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.message || error.message || "Failed to update task",
            status: error.response?.status || 500
        }
    }
}

// Authentication functions
export const login = async (credentials: LoginRequest) => {
    try {
        const response = await apiClient.post<AuthResponse>("/api/v1/auth/login", credentials)
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (login):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.error || error.message || "Login failed",
            status: error.response?.status || 500
        }
    }
}

export const register = async (userData: RegisterRequest) => {
    try {
        const response = await apiClient.post<AuthResponse>("/api/v1/auth/register", userData)
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (register):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.error || error.message || "Registration failed",
            status: error.response?.status || 500
        }
    }
}

export const logout = async () => {
    try {
        const response = await apiClient.post("/api/v1/auth/logout")
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (logout):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.error || error.message || "Logout failed",
            status: error.response?.status || 500
        }
    }
}

export const validateUser = async () => {
    try {
        const response = await apiClient.get<AuthResponse>("/api/v1/auth/validate")
        return {
            success: true,
            data: response.data,
            status: response.status
        }
    } catch (error: any) {
        console.error("API Error (validateUser):", error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.error || error.message || "User validation failed",
            status: error.response?.status || 500
        }
    }
}

export const getUser = async (id?: string) => {
    try {
        if (id) {
            const response = await apiClient.get<{ user: User }>(`/api/v1/user?id=${encodeURIComponent(id)}`)
            return {
                success: true,
                data: response.data.user,
                status: response.status,
            }
        }

        // No id: validate using cookie-based endpoint to get current user
        const response = await apiClient.get<AuthResponse>("/api/v1/auth/validate")
        return {
            success: true,
            data: response.data.user,
            status: response.status,
        }
    } catch (error: any) {
        console.error('API Error (getUser):', error.response?.data || error.message)
        return {
            success: false,
            error: error.response?.data?.error || error.message || 'Failed to fetch user',
            status: error.response?.status || 500,
        }
    }
}

// Export the configured axios instance for other API calls
export { apiClient }