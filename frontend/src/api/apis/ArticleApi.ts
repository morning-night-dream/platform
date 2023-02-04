/* tslint:disable */
/* eslint-disable */
/**
 * Morning Night Dream - AppGateway
 * This is the AppGateway API documentation.
 *
 * The version of the OpenAPI document: 0.0.1
 * Contact: morning.night.dream@example.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


import * as runtime from '../runtime';
import type {
  ListArticleResponse,
} from '../models';
import {
    ListArticleResponseFromJSON,
    ListArticleResponseToJSON,
} from '../models';

export interface V1ListArticlesRequest {
    maxPageSize: number;
    pageToken?: string;
}

/**
 * 
 */
export class ArticleApi extends runtime.BaseAPI {

    /**
     * List articles
     * List articles
     */
    async v1ListArticlesRaw(requestParameters: V1ListArticlesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<Array<ListArticleResponse>>> {
        if (requestParameters.maxPageSize === null || requestParameters.maxPageSize === undefined) {
            throw new runtime.RequiredError('maxPageSize','Required parameter requestParameters.maxPageSize was null or undefined when calling v1ListArticles.');
        }

        const queryParameters: any = {};

        if (requestParameters.pageToken !== undefined) {
            queryParameters['pageToken'] = requestParameters.pageToken;
        }

        if (requestParameters.maxPageSize !== undefined) {
            queryParameters['maxPageSize'] = requestParameters.maxPageSize;
        }

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/article`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.JSONApiResponse(response, (jsonValue) => jsonValue.map(ListArticleResponseFromJSON));
    }

    /**
     * List articles
     * List articles
     */
    async v1ListArticles(requestParameters: V1ListArticlesRequest, initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<Array<ListArticleResponse>> {
        const response = await this.v1ListArticlesRaw(requestParameters, initOverrides);
        return await response.value();
    }

}
