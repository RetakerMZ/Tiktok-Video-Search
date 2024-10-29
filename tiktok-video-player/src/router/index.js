import { createRouter, createWebHistory } from 'vue-router';
import EmbedPlayer from '../components/EmbedPlayer.vue';
import HomePage from "../components/HomePage.vue";

const routes = [
	{
		path: '/embed-player',
		name: 'EmbedPlayer',
		component: EmbedPlayer,
	},
	{
		path: '/',
		name: 'HomePage',
		component: HomePage,
	},
];

const router = createRouter({
	history: createWebHistory(),
	routes,
});

export default router;
