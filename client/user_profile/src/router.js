import Router from "vue-router";
import Vue from "vue";

Vue.use(Router);

// import Dashboard from "Dashboard.vue";

function Path(viewPath) {
  return()=>import(`@/views/${viewPath}.vue`)
}
const routes = [
  {
    path: "/",
    name: "Dashboard",
    component: Path('Dashboard'),
    meta: {
      title: "Dashboard",
    },
  },
  {
    path: "/EditUser/:userId",
    name: "EditUser",
    component: Path('EditUser'),
    meta: {
      title: "Edit User"
    },
  },
  {
    path: "/AddUser",
    name: "AddUser",
    component: Path('AddUser'),
    meta: {
      title: "Add User"
    },
  }
  //   { path: "/bar", component: Bar },
];

export default new Router({
  routes,
});
