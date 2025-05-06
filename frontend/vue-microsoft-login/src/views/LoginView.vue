<template>
  <div class="login-wrapper">
    <div class="sky">
      <div class="cloud cloud-1"></div>
      <div class="cloud cloud-2"></div>
      <div class="stars">
        <div class="star" v-for="i in 10" :key="i" :style="{ left: `${i * 10}%` }"></div>
      </div>
    </div>

  <div class="login-box">
      <img src="/images/bondly-logo.png" class="logo" alt="Bondly Logo" />
      <h1 class="title">Bondly</h1>
      <p class="tagline">Connect. Share. Enjoy.</p>
      <div style="display: flex; justify-content: center;">
        <button class="loginBtn" @click="loginWithMicrosoft">
        <img src="/images/ms.png" alt="Microsoft logo" width="20" height="20" />
        Login with Microsoft
        </button>
      </div>
    </div>
  </div>
</template>

<script setup>
const loginWithMicrosoft = () => {
  window.location.href = 'http://localhost:8000/auth/microsoft';
};
</script>

<style scoped>
@import url('https://fonts.googleapis.com/css2?family=Fredoka+One&display=swap');

.login-wrapper {
  font-family: 'Fredoka One', cursive;
  position: relative;
  height: 98vh;
  background: linear-gradient(to top, #ffe4e1 0%, #add8e6 100%);
  overflow: hidden;
  display: flex;
  align-items: center;
  justify-content: center;
  
}

.sky {
  position: absolute;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  z-index: 0;
}

.cloud {
  background: white;
  border-radius: 50%;
  position: absolute;
  opacity: 0.7;
  animation: float 6s ease-in-out infinite;
}

.cloud-1 {
  width: 200px;
  height: 60px;
  top: 15%;
  left: 20%;
  box-shadow: 60px 0 white, 120px 0 white;
}

.cloud-2 {
  width: 150px;
  height: 50px;
  top: 30%;
  left: 60%;
  box-shadow: 40px 0 white, 80px 0 white;
  animation-delay: 2s;
}

@keyframes float {
  0%, 100% {
    transform: translateY(0px);
  }
  50% {
    transform: translateY(10px);
  }
}

.stars .star {
  position: absolute;
  top: calc(10% + (30 * 1px));
  width: 10px;
  height: 10px;
  background: gold;
  border-radius: 50%;
  animation: twinkle 2s infinite ease-in-out alternate;
}

@keyframes twinkle {
  0% {
    opacity: 0.5;
    transform: scale(1);
  }
  100% {
    opacity: 1;
    transform: scale(1.3);
  }
}

.login-box {
  position: relative;
  z-index: 1;
  background: #fff8dc;
  padding: 2.5rem;
  border-radius: 25px;
  box-shadow: 0 20px 40px rgba(0, 0, 0, 0.2);
  text-align: center;
  width: 350px;
  animation: popIn 0.8s ease forwards;
  transform: scale(0);
}

@keyframes popIn {
  to {
    transform: scale(1);
  }
}

.logo {
  width: 80px;
  margin-bottom: 1rem;
  animation: bounce 0.5s infinite alternate;
}

@keyframes bounce {
  from {
    transform: translateY(0px);
  }
  to {
    transform: translateY(-10px);
  }
}

.title {
  font-size: 2rem;
  color: #ff6b81;
  margin-bottom: 0.5rem;
}

.tagline {
  color: #333;
  font-size: 1rem;
  margin-bottom: 2rem;
}

.loginBtn {
  background-color: #ffe0b2;
  border: 2px solid #000;
  padding: 12px 24px;
  border-radius: 25px;
  font-size: 1rem;
  cursor: pointer;
  display: flex;
  align-items: center;
  gap: 10px;
  justify-content: center;
  transition: all 0.3s ease;
}

.loginBtn:hover {
  background-color: #ffd180;
  transform: scale(1.05);
}


</style>
