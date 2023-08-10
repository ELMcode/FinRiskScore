import { createRouter, createWebHistory } from 'vue-router'
import CalculateRiskScore from '../views/CalculateRiskScore.vue'

const router = createRouter({
  history: createWebHistory(),
  routes: [
    {
      path: '/calculate-risk-score',
      component: CalculateRiskScore
    }
  ]
})

export default router