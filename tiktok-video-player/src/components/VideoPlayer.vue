<template>
  <div class="video-player">
    <video
      ref="videoElement"
      :src="videoUrl"
      controls
      @play="handlePlay"
      @pause="handlePause"
    ></video>
  </div>
</template>

<script>
export default {
  name: "VideoPlayer",
  props: {
    videoUrl: {
      type: String,
      required: true
    }
  },
  methods: {
    handlePlay() {
      console.log("Video playing...");
    },
    handlePause() {
      console.log("Video paused.");
    }
  }
};
</script>

<style scoped>

.video-player {
  max-width: 100%;
}
</style>
