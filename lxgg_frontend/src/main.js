import Vue from 'vue';
import VueRouter from 'vue-router';
import Main from './main.vue';
import Login from './login.vue';
import NewContainer from './new-container.vue';
import Container from './containers.vue';
import Images from './images.vue';
import Network from './network.vue';
import Host from './host.vue';
import Settings from './settings.vue';
import User from './user.vue';

Vue.use(VueRouter);

const checkLoggedIn = (to, from, next) => {
	//redirect to login if no cookie found
	if(document.cookie.indexOf('lxgg_session=') == -1) {
		next('/login');
	} else {
		next();
	}
}

const routes = [
	{ path: '/login', name: 'login', component: Login},
	{ path: '/new', name: 'new', component: NewContainer, beforeEnter: checkLoggedIn},
	{ path: '/', alias: '/containers', name: 'containers', component: Container, beforeEnter: checkLoggedIn},
	{ path: '/images', name: 'images', component: Images, beforeEnter: checkLoggedIn},
	{ path: '/networking', name: 'networking', component: Network, beforeEnter: checkLoggedIn},
	{ path: '/host', name: 'host', component: Host, beforeEnter: checkLoggedIn},
	{ path: '/settings', name: 'settings', component: Settings, beforeEnter: checkLoggedIn},
	{ path: '/user', name: 'user', component: User, beforeEnter: checkLoggedIn},
];

const router = new VueRouter({ routes });

const app = new Vue({
	el: '#app',
	router: router,
	render: h => h(Main, {
		// props: {
		// 	greeting: "Hello"
		// }
	})
});