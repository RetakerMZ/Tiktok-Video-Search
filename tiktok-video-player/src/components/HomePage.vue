<template>
  <div id="home">
		<h1>TikTok Video Search</h1>
    <SearchBar />
  </div>
</template>

<script>
import SearchBar from "@/components/SearchBar.vue";

export default {
  components: { SearchBar },
};
</script>

<style>
#home {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>