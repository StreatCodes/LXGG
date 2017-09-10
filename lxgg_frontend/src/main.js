import Main from './main.vue'
import NewContainer from './new-container.vue'
import Container from './containers.vue'
import Images from './images.vue'
import Network from './network.vue'
import Host from './host.vue'
import Settings from './settings.vue'
import User from './user.vue'


const routes = [
	{ path: '/new', name: 'new', component: NewContainer},
	{ path: '/', alias: '/containers', name: 'containers', component: Container},
	{ path: '/images', name: 'images', component: Images},
	{ path: '/networking', name: 'networking', component: Network},
	{ path: '/host', name: 'host', component: Host},
	{ path: '/settings', name: 'settings', component: Settings},
	{ path: '/user', name: 'user', component: User},
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