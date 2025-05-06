<template>
    <div class="post-form">
      <form @submit.prevent="submitPost">
        <textarea v-model="content" placeholder="What's on your mind?" required></textarea>
        <input type="file" @change="onFileChange" />
        <button type="submit">Post</button>
      </form>
    </div>
  </template>
  
  <script setup>
  import { ref } from 'vue'
  
  const content = ref('')
  const image = ref(null)
  
  const onFileChange = (e) => {
    image.value = e.target.files[0]
  }
  
  const submitPost = async () => {
    const formData = new FormData()
    formData.append('content', content.value)
    if (image.value) formData.append('image', image.value)
  
    try {
      const res = await fetch('http://localhost:8000/posts', {
        method: 'POST',
        body: formData,
      })
      const data = await res.json()
      alert('Post created!')
      content.value = ''
      image.value = null
    } catch (err) {
      console.error(err)
      alert('Failed to create post')
    }
  }
  </script>
  
  <style scoped>
  .post-form {
    margin-bottom: 1rem;
  }
  textarea {
    width: 100%;
    min-height: 100px;
  }
  </style>
  