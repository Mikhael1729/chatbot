import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";
import PrimeVue from "primevue/config";
import Button from "primevue/button";
import InputText from "primevue/inputtext";

import "primevue/resources/themes/saga-blue/theme.css"; //theme
import "primevue/resources/primevue.min.css"; //core css
import "primeicons/primeicons.css"; //icons

createApp(App)
  // Plugins.
  .use(store)
  .use(router)
  .use(PrimeVue)
  // PrimeVue components
  .component("Button", Button)
  .component("InputText", InputText)
  // Mounting
  .mount("#app");

