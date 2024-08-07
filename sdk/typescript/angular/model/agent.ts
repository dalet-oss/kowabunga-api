/**
 * Kowabunga API documentation
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * The version of the OpenAPI document: 0.42.0
 * Contact: csops@dalet.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * A Kowabunga remote agent.
 */
export interface Agent { 
    /**
     * The Kowabunga remote agent ID (auto-generated).
     */
    id?: string;
    /**
     * The Kowabunga remote agent name.
     */
    name: string;
    /**
     * The Kowabunga remote agent description.
     */
    description?: string;
    /**
     * The Kowabunga agent type.
     */
    type: Agent.TypeEnum;
}
export namespace Agent {
    export type TypeEnum = 'KCA' | 'KSA' | 'KNA';
    export const TypeEnum = {
        Kca: 'KCA' as TypeEnum,
        Ksa: 'KSA' as TypeEnum,
        Kna: 'KNA' as TypeEnum
    };
}


