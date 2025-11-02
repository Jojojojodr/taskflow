<template>
    <div class="max-w-lg mx-auto mt-10 p-6 bg-white rounded-lg shadow-md">
        <h2 class="text-2xl font-bold mb-4">Profile</h2>
        <div v-if="loading" class="text-gray-500">Loading...</div>
        <div v-else-if="user">
            <div class="mb-2">
                <span class="font-semibold">Name:</span> {{ user.name }}
            </div>
            <div class="mb-2">
                <span class="font-semibold">Email:</span> {{ user.email || 'N/A' }}
            </div>
            <div class="mb-2">
                <span class="font-semibold">ID:</span> {{ user.id }}
            </div>
        </div>
        <div v-else class="text-red-500">You are not logged in.</div>
    </div>
</template>

<script setup>
import { ref, onMounted } from 'vue'
import { getUser } from '../utils/api'

const user = ref(null)
const loading = ref(true)

const fetchUser = async () =>
{
    loading.value = true
    const res = await getUser()
    if (res.success && res.data) {
        user.value = res.data
    } else {
        user.value = null
    }
    loading.value = false
}

onMounted(fetchUser)
</script>
