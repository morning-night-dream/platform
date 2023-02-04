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

/**
 * 
 */
export class HealthApi extends runtime.BaseAPI {

    /**
     * ヘルスチェック
     * ヘルスチェック
     */
    async v1HealthRaw(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<runtime.ApiResponse<void>> {
        const queryParameters: any = {};

        const headerParameters: runtime.HTTPHeaders = {};

        const response = await this.request({
            path: `/v1/health`,
            method: 'GET',
            headers: headerParameters,
            query: queryParameters,
        }, initOverrides);

        return new runtime.VoidApiResponse(response);
    }

    /**
     * ヘルスチェック
     * ヘルスチェック
     */
    async v1Health(initOverrides?: RequestInit | runtime.InitOverrideFunction): Promise<void> {
        await this.v1HealthRaw(initOverrides);
    }

}