# KawaiiIpSec
## Properties

| Name | Type | Description | Notes |
|------------ | ------------- | ------------- | -------------|
| **id** | **String** | The Kawaii IPsec connection ID (auto-generated). | [optional] [default to null] |
| **name** | **String** | The Kawaii IPsec connection name. | [default to null] |
| **description** | **String** | The Kawaii IPsec connection description. | [optional] [default to null] |
| **ip** | **String** | The Kawaii IPsec connection IPSec IP. | [optional] [default to null] |
| **remote\_ip** | **String** | The Kawaii IPsec connection remote peer VPN Gateway. | [default to null] |
| **remote\_subnet** | **String** | The Kawaii IPsec connection remote subnet. | [default to null] |
| **pre\_shared\_key** | **String** | The Kawaii IPsec connection pre-shared key(PSK). | [default to null] |
| **dpd\_timeout\_action** | **String** | The Kawaii IPsec connection Dead Peer Detection Action (clear,restart or trap). | [optional] [default to restart] |
| **dpd\_timeout** | **String** | The Kawaii IPsec connection Dead Peer Detection Timeout. | [optional] [default to 240s] |
| **start\_action** | **String** | The Kawaii IPsec connection start action (none, start, trap). | [optional] [default to start] |
| **rekey\_time** | **String** | The Kawaii IPsec connection rekey time. Default is 2h. | [optional] [default to 2h] |
| **phase1\_lifetime** | **String** | The Kawaii IPsec connection Lifetime for phase 1 negociation. Default is 1h. | [optional] [default to 1h] |
| **phase1\_dh\_group\_number** | **Long** | The Kawaii IPsec connection phase 1 Diffie Hellman IANA algorithm. | [optional] [default to null] |
| **phase1\_integrity\_algorithm** | **String** | The Kawaii IPsec connection phase 1 integrity algorithm.. | [optional] [default to null] |
| **phase1\_encryption\_algorithm** | **String** | The Kawaii IPsec connection phase 1 encryption algorithm.. | [optional] [default to null] |
| **phase2\_lifetime** | **String** | The Kawaii IPsec connection Lifetime for phase 2 negociation. Default is 1h. | [optional] [default to 1h] |
| **phase2\_dh\_group\_number** | **Long** | The Kawaii IPsec connection phase 2 Diffie Hellman IANA algorithm. | [optional] [default to null] |
| **phase2\_integrity\_algorithm** | **String** | The Kawaii IPsec connection phase 2 integrity algorithm.. | [optional] [default to null] |
| **phase2\_encryption\_algorithm** | **String** | The Kawaii IPsec connection phase 2 encryption algorithm.. | [optional] [default to null] |
| **firewall** | [**KawaiiFirewall**](.md) | The Kawaii IPsec connection firewall settings on the IPsec interface. | [optional] [default to null] |

[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)

