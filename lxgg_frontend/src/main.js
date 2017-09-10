import Main from './main.vue'
import NewContainer from './new-container.vue'
import Container from './containers.vue'


const routes = [
	{ path: '/new', component: NewContainer},
	{ path: '/', component: Container},
	{ path: '/containers', component: Container},
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