// src/router/index.ts

import { createRouter, createWebHistory } from "vue-router";
import type {RouteRecordRaw} from 'vue-router';
import Home from "@/views/Home.vue";
import CalculateRiskScore from "@/views/CalculateRiskScore.vue";

const routes: Array<RouteRecordRaw> = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/calculate-risk-score",
    name: "CalculateRiskScore",
    component: CalculateRiskScore,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
