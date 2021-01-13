/* tslint:disable */
/* eslint-disable */
/**
 * SUT SA Example API
 * This is a sample server for SUT SE 2563
 *
 * The version of the OpenAPI document: 1.0
 * Contact: support@swagger.io
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */

import { exists, mapValues } from '../runtime';
import {
    EntBloodTypeEdges,
    EntBloodTypeEdgesFromJSON,
    EntBloodTypeEdgesFromJSONTyped,
    EntBloodTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntBloodType
 */
export interface EntBloodType {
    /**
     * Blood holds the value of the "blood" field.
     * @type {string}
     * @memberof EntBloodType
     */
    blood?: string;
    /**
     * 
     * @type {EntBloodTypeEdges}
     * @memberof EntBloodType
     */
    edges?: EntBloodTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntBloodType
     */
    id?: number;
}

export function EntBloodTypeFromJSON(json: any): EntBloodType {
    return EntBloodTypeFromJSONTyped(json, false);
}

export function EntBloodTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntBloodType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'blood': !exists(json, 'blood') ? undefined : json['blood'],
        'edges': !exists(json, 'edges') ? undefined : EntBloodTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntBloodTypeToJSON(value?: EntBloodType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'blood': value.blood,
        'edges': EntBloodTypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}


