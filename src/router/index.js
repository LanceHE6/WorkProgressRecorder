import { createRouter, createWebHashHistory } from 'vue-router';
import Main from '../views/Main.vue';


const routes = [
    { path: '/', component: Main , name: 'main'},
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export {router, routes};
