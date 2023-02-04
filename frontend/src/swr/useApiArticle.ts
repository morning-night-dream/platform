import useSWR, { useSWRConfig } from 'swr';

import { V1ListArticlesRequest } from '../api/apis/ArticleApi';
import { Article } from '../api/models';
import { client } from './client';

// 1回に取得する記事の数
const articlesPerPage = 20;


type ArticlesState = {
    data: Article[];
    error?: Error;
    currentIndex?: string;
};

const articlesState: ArticlesState = {
    data: [],
};

export const useListArticles = () => {
    const key = `/api/v1/articles`;

    const request : V1ListArticlesRequest = {
        maxPageSize: articlesPerPage,
        pageToken: articlesState.currentIndex ?? "",
    };

    const fetcher = async () =>  client.v1ListArticles(request);
    const { data } = useSWR(key, fetcher);

    const fetchedArticles = data?.articles ?? [];
    const existIds = articlesState.data.map(d => d.id);
    const additionalArticles = fetchedArticles.filter(d => existIds.indexOf(d.id) < 0);
    articlesState.data.push(...additionalArticles);

    // NextPageTokenが空の場合、もうこれ以上データがないのでcurrentIndexを更新しない
    if (!(data?.nextPageToken === '')) {
        articlesState.currentIndex = data?.nextPageToken;
    }

    const { mutate } = useSWRConfig();

    return {
        data: articlesState.data,
        async mutate() {
            await mutate(key);
        },
    };
};
