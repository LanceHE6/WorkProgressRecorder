import { createRouter, createWebHashHistory } from 'vue-router';
import Main from '../views/Main.vue';
import Postgraduate from "../views/Postgraduate.vue";
import Occupation from "../views/Occupation.vue";
import UserCenter from "../views/UserCenter.vue";
import LogDetail from "../views/LogDetail.vue";
import UserManagement from "../views/UserManagement.vue";

const routes = [
    { path: '/', component: Main , name: 'main', meta:{index:0},},
    { path: '/postgraduate', component: Postgraduate , name: 'postgraduate', meta:{index:1},},
    { path: '/occupation', component: Occupation , name: 'occupation', meta:{index:1},},
    { path: '/logDetail', component: LogDetail , name: 'logDetail', meta:{index:2},},
    { path: '/userCenter', component: UserCenter , name: 'userCenter', meta:{index:1},},
    { path: '/userManagement', component: UserManagement , name: 'userManagement', meta:{index:1},},
];

const router = createRouter({
    history: createWebHashHistory(),
    routes,
});

export {router, routes};
