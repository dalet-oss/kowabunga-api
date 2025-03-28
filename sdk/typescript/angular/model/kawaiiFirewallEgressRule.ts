/**
 * Kowabunga API documentation
 *
 * Contact: maintainers@kowabunga.cloud
 *
 * NOTE: This class is auto generated by OpenAPI Generator (https://openapi-generator.tech).
 * https://openapi-generator.tech
 * Do not edit the class manually.
 */


/**
 * A Kawaii public firewall egress rule.
 */
export interface KawaiiFirewallEgressRule { 
    /**
     * The destination IP or CIDR to accept/drop public traffic to.
     */
    destination?: string;
    /**
     * The transport layer protocol to accept/drop public traffic to.
     */
    protocol?: KawaiiFirewallEgressRule.ProtocolEnum;
    /**
     * The port (or list of ports) to accept/drop public traffic from. Ranges are accepted. Format is a-b,c-d (e.g. 443; 22,80,443; 80,443,3000-3005).
     */
    ports: string;
}
export namespace KawaiiFirewallEgressRule {
    export type ProtocolEnum = 'tcp' | 'udp';
    export const ProtocolEnum = {
        Tcp: 'tcp' as ProtocolEnum,
        Udp: 'udp' as ProtocolEnum
    };
}


