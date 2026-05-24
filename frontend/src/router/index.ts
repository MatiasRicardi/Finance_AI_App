import { createRouter, createWebHistory } from '@ionic/vue-router'
import type { RouteRecordRaw } from 'vue-router'
import { useAuthStore } from '@/stores/auth.store'

const routes: RouteRecordRaw[] = [
  {
    path: '/',
    redirect: '/app/dashboard',
  },
  {
    path: '/auth',
    children: [
      {
        path: 'login',
        name: 'Login',
        component: () => import('@/pages/auth/LoginPage.vue'),
      },
      {
        path: 'register',
        name: 'Register',
        component: () => import('@/pages/auth/RegisterPage.vue'),
      },
    ],
  },
  {
    path: '/app',
    meta: { requiresAuth: true },
    component: () => import('@/pages/AppTabs.vue'),
    children: [
      {
        path: 'dashboard',
        name: 'Dashboard',
        component: () => import('@/pages/DashboardPage.vue'),
      },
      {
        path: 'transactions',
        name: 'Transactions',
        component: () => import('@/pages/TransactionsPage.vue'),
      },
      {
        path: 'import',
        name: 'Import',
        component: () => import('@/pages/ImportPage.vue'),
      },
      {
        path: 'import/:id/review',
        name: 'ImportReview',
        component: () => import('@/pages/ImportReviewPage.vue'),
      },
      {
        path: 'import/:id/result',
        name: 'ImportResult',
        component: () => import('@/pages/ImportResultPage.vue'),
      },
      {
        path: 'reports',
        name: 'Reports',
        component: () => import('@/pages/ReportsPage.vue'),
      },
      {
        path: 'profile',
        name: 'Profile',
        component: () => import('@/pages/ProfilePage.vue'),
      },
      {
        path: 'settings/ai-provider',
        name: 'AIProviderSettings',
        component: () => import('@/pages/settings/AIProviderSettingsPage.vue'),
      },
      {
        path: 'settings/categories',
        name: 'Categories',
        component: () => import('@/pages/settings/CategoriesPage.vue'),
      },
    ],
  },
  {
    path: '/:pathMatch(.*)*',
    redirect: '/app/dashboard',
  },
]

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
})

router.beforeEach((to) => {
  const authStore = useAuthStore()

  if (to.meta.requiresAuth && !authStore.isAuthenticated) {
    return { name: 'Login' }
  }

  if ((to.name === 'Login' || to.name === 'Register') && authStore.isAuthenticated) {
    return { name: 'Dashboard' }
  }
})

export default router
