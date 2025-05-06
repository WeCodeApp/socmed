<template>
    <div class="flex items-center justify-center h-screen">
      <p v-if="isLoading">Logging you in...</p>
      <p v-if="error">An error occurred while logging in.</p>
    </div>
  </template>
  
  <script setup>
  import { onMounted, ref } from 'vue'
  import { useRouter, useRoute } from 'vue-router'
  
  const router = useRouter()
  const route = useRoute()
  
  const isLoading = ref(true)
  const error = ref(false)
  
  onMounted(async () => {
    try {
      const token = route.query.token
        
      if (token) {
        // Save the token to local storage or state management
        localStorage.setItem('auth_token', token)
  
        // Redirect to the home page after successful login
        router.push({ name: 'Home' })
      } else {
        // Handle errors (if no token is provided)
        error.value = true
      }
} catch (err) {
      console.error(err)
      error.value = true
    } finally {
      isLoading.value = false
    }   
  })
  </script>
  