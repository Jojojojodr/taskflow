import { createRouter, createWebHistory } from 'vue-router'
import Home from '../views/Home.vue'
import Tasks from '../views/Tasks.vue'
import Board from '../views/Board.vue'
import Profile from '../views/Profile.vue'

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/tasks',
        name: 'Tasks',
        component: Tasks
    },
    {
        path: '/board',
        name: 'Board',
        component: Board
    },
    {
        path: '/profile',
        name: 'Profile',
        component: Profile
    }
]

const router = createRouter({
    history: createWebHistory(),
    routes
})

export default router