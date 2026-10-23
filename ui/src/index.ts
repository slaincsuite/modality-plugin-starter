import { createApp } from 'vue'
import HomeView from './views/HomeView.vue'

export function mount(selector: string | Element = '#app') {
  return createApp(HomeView).mount(selector)
}

export default {
  routes: [],
  extensionPoints: {},
}
