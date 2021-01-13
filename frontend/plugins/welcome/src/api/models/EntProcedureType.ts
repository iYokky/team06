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
    EntProcedureTypeEdges,
    EntProcedureTypeEdgesFromJSON,
    EntProcedureTypeEdgesFromJSONTyped,
    EntProcedureTypeEdgesToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntProcedureType
 */
export interface EntProcedureType {
    /**
     * ProcedureName holds the value of the "ProcedureName" field.
     * @type {string}
     * @memberof EntProcedureType
     */
    procedureName?: string;
    /**
     * 
     * @type {EntProcedureTypeEdges}
     * @memberof EntProcedureType
     */
    edges?: EntProcedureTypeEdges;
    /**
     * ID of the ent.
     * @type {number}
     * @memberof EntProcedureType
     */
    id?: number;
}

export function EntProcedureTypeFromJSON(json: any): EntProcedureType {
    return EntProcedureTypeFromJSONTyped(json, false);
}

export function EntProcedureTypeFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntProcedureType {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'procedureName': !exists(json, 'ProcedureName') ? undefined : json['ProcedureName'],
        'edges': !exists(json, 'edges') ? undefined : EntProcedureTypeEdgesFromJSON(json['edges']),
        'id': !exists(json, 'id') ? undefined : json['id'],
    };
}

export function EntProcedureTypeToJSON(value?: EntProcedureType | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'ProcedureName': value.procedureName,
        'edges': EntProcedureTypeEdgesToJSON(value.edges),
        'id': value.id,
    };
}

