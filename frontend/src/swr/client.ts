import { ArticleApi } from '../api/apis/ArticleApi';
import { Configuration } from '../api';

export const client = new ArticleApi(
    new Configuration({
        basePath: (import.meta.env.VITE_SERVICE_ENDPOINT as string) + '/api',
        headers: {
            'Content-Type': 'application/json',
        },
    })
);
