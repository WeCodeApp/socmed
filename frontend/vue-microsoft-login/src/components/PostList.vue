<template>
    <div>
      <h2 class="text-xl mb-2">Posts</h2>
      <div v-if="posts.length === 0" class="text-gray-500">No posts available.</div>
      <div v-for="post in posts" :key="post.id" class="border p-4 mb-4 rounded">
        <div class="mb-2 text-gray-700">{{ post.content }}</div>
        <img v-if="post.image_url" :src="post.image_url" class="max-w-full mb-2 rounded" />
        <div class="text-sm text-gray-500">Posted by: {{ post.user?.display_name || 'Anonymous' }}</div>
        <div class="mt-2 flex gap-2">
          <button @click="$emit('edit', post)" class="text-blue-500">Edit</button>
          <button @click="$emit('delete', post.id)" class="text-red-500">Delete</button>
        </div>
      </div>
    </div>
  </template>
  
  <script setup>
  defineProps({
    posts: Array
  })
  
  defineEmits(['edit', 'delete'])
  </script>
  