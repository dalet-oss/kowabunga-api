/**
 * Kowabunga API documentation
 * Kvm Orchestrator With A BUNch of Goods Added
 *
 * The version of the OpenAPI document: 0.37.0
 * Contact: csops@dalet.com
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * A KGW public firewall egress rule.
 */
export interface KGWFirewallEgressRule { 
    /**
     * The destination IP or CIDR to accept/drop public traffic to.
     */
    destination?: string;
    /**
     * The transport layer protocol to accept/drop public traffic to.
     */
    protocol?: KGWFirewallEgressRule.ProtocolEnum;
    /**
     * The port (or list of ports) to accept/drop public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005).
     */
    ports: string;
}
export namespace KGWFirewallEgressRule {
    export type ProtocolEnum = 'tcp' | 'udp';
    export const ProtocolEnum = {
        Tcp: 'tcp' as ProtocolEnum,
        Udp: 'udp' as ProtocolEnum
    };
}

