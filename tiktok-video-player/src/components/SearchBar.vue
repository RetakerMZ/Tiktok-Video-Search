<template>
  <div>
    <input
      type="text"
      v-model="searchQuery"
      placeholder="Search for TikTok videos"
      @keyup.enter="performInitialSearch"
    />
    <button @click="performInitialSearch">Search</button>

		<div class="video-list">
      <div
        class="video-card"
        v-for="(video, index) in videos"
        :key="index"
        @click="openModal(video.play)"
      >
        <img :src="video.cover" alt="Thumbnail" class="thumbnail" />
        <h3 class="video-title">{{ video.title }}</h3>

				<div class="author-info">
          <img :src="video.author.avatar" alt="Author Avatar" class="avatar" />
          <p class="video-author">{{ video.author.nickname }}</p>
        </div>
      </div>
    </div>

		<div v-if="isLoading" class="loading-spinner">
      <img src="../assets/loading.gif" alt="Loading..." />
    </div>

    <VideoModal
      v-if="isModalOpen"
      :videoUrl="embedVideoUrl"
      @close="isModalOpen = false"
    />
  </div>
</template>

<script>
import axios from "axios";
import VideoModal from './VideoModal.vue';

export default {
  components: { VideoModal },
	props: ["video"],
  data() {
    return {
			searchQuery: "",
      videos: [],
      selectedVideoUrl: null,
      cursor: 0,
      isLoading: false,
			seenUrls: new Set(),
			isModalOpen: false,
    };
  },
  methods: {
    performInitialSearch() {
      if (this.searchQuery && !this.isLoading ) {
				this.isLoading = true;
				this.cursor = 0;
        axios
          .get(
            `${process.env.VUE_APP_API_BASE_URL}/search`,
            {
              params: { query: this.searchQuery, cursor: this.cursor },
            }
					)
          .then((response) => {
            this.videos = response.data.data.data.videos.map(video => ({
              title: video.title,
              play: video.play,
              cover: video.cover,
              author: video.author,
            }));
						this.isLoading = false;
          })
          .catch((error) => {
            console.error("Error fetching search results:", error);
          });
      }
    },
		async performSearch() {
      if (this.searchQuery && !this.isLoading) {
        this.isLoading = true;
        try {
          const response = await axios.get(
            `${process.env.VUE_APP_API_BASE_URL}/search`,
            {
              params: { query: this.searchQuery, cursor: this.cursor },
            }
					);
          const newVideos = response.data.data.data.videos.map(video => ({
            title: video.title,
            play: video.play,
            cover: video.cover,
            author: video.author,
          }));
					const uniqueVideos = newVideos.filter(video => {
              if (!this.seenUrls.has(video.play)) {
                this.seenUrls.add(video.play);
                return true;
              }
              return false;
          });
          this.videos = [...this.videos, ...uniqueVideos];
					if (newVideos.length > 0) {
            this.cursor += 30;
          }
        } catch (error) {
          console.error("Error fetching search results:", error);
        } finally {
          this.isLoading = false;
        }
      }
    },
    openModal(videoUrl) {
			this.embedVideoUrl = videoUrl;
      this.isModalOpen = true;
    },
		handleScroll() {
      const bottomReached =
        window.innerHeight + window.scrollY >= document.documentElement.scrollHeight;
      if (bottomReached && !this.isLoading) {
        this.performSearch(); // Load more videos when reaching the bottom
      }
    }
  },
	mounted() {
    window.addEventListener("scroll", this.handleScroll);
  },
  beforeUnmount() {
    window.removeEventListener("scroll", this.handleScroll);
  }
};
</script>

<style scoped>

input {
  padding: 8px;
  margin-right: 8px;
}

button {
  padding: 8px;
  background-color: #42b983;
  color: white;
  border: none;
  cursor: pointer;
}

button:hover {
  background-color: #38a169;
}

.video-list {
  display: flex;
  flex-wrap: wrap;
  justify-content: space-around;
  margin: 20px 0;
}

.video-card {
  background-color: #f9f9f9;
  border: 1px solid #ddd;
  border-radius: 5px;
  margin: 10px;
  width: 200px;
  cursor: pointer;
  transition: box-shadow 0.3s;
  display: flex;
  flex-direction: column;
  align-items: center;
	overflow: hidden;
}

.video-card:hover {
  box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.thumbnail {
  width: 100%;
  height: auto;
  border-radius: 5px;
	border-radius: 0;
  display: block;
}

.video-title {
  font-size: 0.75rem;
  font-weight: bold;
  margin: 10px 0px 5px;
  text-align: center;
  white-space: normal;
  overflow: hidden;
  text-overflow: ellipsis;
	width:100%;
	display: -webkit-box;
	-webkit-line-clamp: 3; 
  -webkit-box-orient: vertical;
	line-height: 1.2em; /* Adjust based on font size */
  max-height: 3.6em; /* 3 lines * line-height */
}

.author-info {
  display: flex;
  align-items: center;
  justify-content: flex-start;
  width: 100%;
	margin-left: 5px;
  margin-top: 5px;
  margin-bottom: 8px;
}

.avatar {
  width: 30px;
  height: 30px;
  border-radius: 50%;
  margin-right: 8px;
}

.video-author {
  font-size: 0.85rem;
  color: #666;
}

.loading-spinner {
  display: flex;
  justify-content: center;
  align-items: center;
  margin: 20px 0;
}
.loading-spinner img {
  width: 50px;
}
</style>
