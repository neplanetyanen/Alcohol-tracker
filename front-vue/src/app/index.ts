import { createApp } from 'vue';

import App from './index.vue';
import { router } from './providers/router';
import { store } from './providers/store';

export const app = createApp(App).use(store).use(router);
