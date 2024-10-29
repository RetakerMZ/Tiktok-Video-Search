<template>
  <div class="video-player-container">
		<video ref="playerDiv"
      class="mirax-player"
			:data-player-theme="themeColor"
      :src="videoUrl">
    </video>
  </div>
</template>

<script>
import { miraxPlayer } from 'mirax-player';
import { useRoute } from "vue-router";
import { ref, onMounted } from "vue";

export default {
	name: "EmbedPlayer",
  setup() {
    const route = useRoute();
    
    const videoUrl = "https://file-examples.com/wp-content/storage/2017/04/file_example_MP4_480_1_5MG.mp4";
    const themeColor = route.query.color || '#000000'; 

		const playerDiv = ref(null);
		console.log(playerDiv.value)
    onMounted(() => {
        miraxPlayer(playerDiv.value);
    });

    return { videoUrl, themeColor, playerDiv };
  },
};
</script>

<style scoped>
.video-player-container {
  width: 100%;
  height: 100vh;
}
</style>
