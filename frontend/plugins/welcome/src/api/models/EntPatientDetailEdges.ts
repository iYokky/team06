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
    EntBloodType,
    EntBloodTypeFromJSON,
    EntBloodTypeFromJSONTyped,
    EntBloodTypeToJSON,
    EntGender,
    EntGenderFromJSON,
    EntGenderFromJSONTyped,
    EntGenderToJSON,
    EntPatient,
    EntPatientFromJSON,
    EntPatientFromJSONTyped,
    EntPatientToJSON,
    EntPrefix,
    EntPrefixFromJSON,
    EntPrefixFromJSONTyped,
    EntPrefixToJSON,
} from './';

/**
 * 
 * @export
 * @interface EntPatientDetailEdges
 */
export interface EntPatientDetailEdges {
    /**
     * 
     * @type {EntBloodType}
     * @memberof EntPatientDetailEdges
     */
    bloodtype?: EntBloodType;
    /**
     * 
     * @type {EntGender}
     * @memberof EntPatientDetailEdges
     */
    gender?: EntGender;
    /**
     * 
     * @type {EntPatient}
     * @memberof EntPatientDetailEdges
     */
    patient?: EntPatient;
    /**
     * 
     * @type {EntPrefix}
     * @memberof EntPatientDetailEdges
     */
    prefix?: EntPrefix;
}

export function EntPatientDetailEdgesFromJSON(json: any): EntPatientDetailEdges {
    return EntPatientDetailEdgesFromJSONTyped(json, false);
}

export function EntPatientDetailEdgesFromJSONTyped(json: any, ignoreDiscriminator: boolean): EntPatientDetailEdges {
    if ((json === undefined) || (json === null)) {
        return json;
    }
    return {
        
        'bloodtype': !exists(json, 'Bloodtype') ? undefined : EntBloodTypeFromJSON(json['Bloodtype']),
        'gender': !exists(json, 'Gender') ? undefined : EntGenderFromJSON(json['Gender']),
        'patient': !exists(json, 'Patient') ? undefined : EntPatientFromJSON(json['Patient']),
        'prefix': !exists(json, 'Prefix') ? undefined : EntPrefixFromJSON(json['Prefix']),
    };
}

export function EntPatientDetailEdgesToJSON(value?: EntPatientDetailEdges | null): any {
    if (value === undefined) {
        return undefined;
    }
    if (value === null) {
        return null;
    }
    return {
        
        'bloodtype': EntBloodTypeToJSON(value.bloodtype),
        'gender': EntGenderToJSON(value.gender),
        'patient': EntPatientToJSON(value.patient),
        'prefix': EntPrefixToJSON(value.prefix),
    };
}

