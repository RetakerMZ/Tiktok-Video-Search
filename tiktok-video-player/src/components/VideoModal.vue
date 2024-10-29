<template>
  <div class="modal-overlay" @click.self="closeModal">
    <div class="modal-content">
      <button class="close-button" @click="closeModal">X</button>
			<div v-if="isLoading" class="loading-spinner">
        <img src="../assets/loading.gif" alt="Loading video..." />
      </div>
      <iframe
        :src="videoUrl"
        width="100%"
        height="100%"
        frameborder="0"
        allow="autoplay; encrypted-media"
        @load="onVideoLoad"
      ></iframe>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    videoUrl: {
      type: String,
      required: true
    }
  },
	data() {
    return {
      isLoading: true 
    };
  },
  methods: {
    closeModal() {
      this.$emit('close');
    },
		onVideoLoad() {
      this.isLoading = false; 
    }
  }
};
</script>

<style scoped>
.modal-overlay {
  position: fixed;
  top: 0;
  left: 0;
  width: 100vw;
  height: 100vh;
  background: rgba(0, 0, 0, 0.5);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 1000;
}

.modal-content {
  background: white;
  padding: 0;
  width: 80%;
  max-width: 800px;
  height: 80vh;
  position: relative;
  border-radius: 8px;
  overflow: hidden;
}

.close-button {
  position: absolute;
  top: 10px;
  right: 10px;
  background: transparent;
  border: none;
  font-size: 1.5rem;
  cursor: pointer;
  color: #333;
}

.loading-spinner {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%; 
}

.loading-spinner img {
  width: 50px; 
}
</style>
