package advancedsecurity

import (
	"context"
	"encoding/xml"
	"time"

	"github.com/videonext/onvif/soap"
)

// against "unused imports"
var _ time.Time
var _ xml.Name

// Unique identifier for keys in the keystore.

type KeyID NCName

// Unique identifier for certificates in the keystore.

type CertificateID NCName

// Unique identifier for certification paths in the keystore.

type CertificationPathID NCName

// Unique identifier for passphrases in the keystore.

type PassphraseID NCName

// Unique identifier for 802.1X configurations in the keystore.

type Dot1XID NCName

// The status of a key in the keystore.

type KeyStatus string

const (

	// Key is ready for use
	KeyStatusOk KeyStatus = "ok"

	// Key is being generated
	KeyStatusGenerating KeyStatus = "generating"

	// Key has not been successfully generated and cannot be used.
	KeyStatusCorrupt KeyStatus = "corrupt"
)

// An object identifier (OID) in dot-decimal form as specified in RFC4512.

type DotDecimalOID string

// The distinguished name attribute type encoded as specified in RFC 4514.

type DNAttributeType string

type DNAttributeValue string

// A base64-encoded ASN.1 value.

type Base64DERencodedASN1Value []byte

// A list of supported 802.1X authentication methods, such as "EAP-PEAP/MSCHAPv2" and "EAP-MD5".  The '/' character is used as a separator between the outer and inner methods.

type Dot1XMethods []string

type CRLID NCName

type CertPathValidationPolicyID NCName

// A list of RSA key lenghts in bits.

type RSAKeyLengths []NonNegativeInteger

// A list of X.509 versions.

type X509Versions []int32

// A list of TLS versions.

type TLSVersions []string

// A list of password based encryption algorithms.

type PasswordBasedEncryptionAlgorithms []string

// A list of password based MAC algorithms.

type PasswordBasedMACAlgorithms []string

type GetServiceCapabilities struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetServiceCapabilities"`
}

type GetServiceCapabilitiesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetServiceCapabilitiesResponse"`

	// The capabilities for the security configuration service is returned in the Capabilities element.
	Capabilities Capabilities `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Capabilities,omitempty"`
}

type CreateRSAKeyPair struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateRSAKeyPair"`

	// The length of the key to be created.
	KeyLength NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema KeyLength,omitempty"`

	// The client-defined alias of the key.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`
}

type CreateRSAKeyPairResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateRSAKeyPairResponse"`

	// The key ID of the key pair being generated.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`

	// Best-effort estimate of how long the key generation will take.
	EstimatedCreationTime Duration `xml:"http://www.onvif.org/ver10/schema EstimatedCreationTime,omitempty"`
}

type UploadKeyPairInPKCS8 struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadKeyPairInPKCS8"`

	// The key pair to be uploaded in a PKCS#8 data structure.
	KeyPair Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyPair,omitempty"`

	// The client-defined alias of the key pair.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// The ID of the passphrase to use for decrypting the uploaded key pair.
	EncryptionPassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl EncryptionPassphraseID,omitempty"`

	// The passphrase to use for decrypting the uploaded key pair.
	EncryptionPassphrase string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl EncryptionPassphrase,omitempty"`
}

type UploadKeyPairInPKCS8Response struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadKeyPairInPKCS8Response"`

	// The key ID of the uploaded key pair.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type UploadCertificateWithPrivateKeyInPKCS12 struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCertificateWithPrivateKeyInPKCS12"`

	// The certificates and key pair to be uploaded in a PKCS#12 data structure.
	CertWithPrivateKey Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertWithPrivateKey,omitempty"`

	// The client-defined alias of the certification path.
	CertificationPathAlias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathAlias,omitempty"`

	// The client-defined alias of the key pair.
	KeyAlias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyAlias,omitempty"`

	// True if and only if the device shall behave as if the client had only supplied the first certificate in the sequence of certificates.
	IgnoreAdditionalCertificates bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl IgnoreAdditionalCertificates,omitempty"`

	// The ID of the passphrase to use for integrity checking of the uploaded PKCS#12 data structure.
	IntegrityPassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl IntegrityPassphraseID,omitempty"`

	// The ID of the passphrase to use for decrypting the uploaded PKCS#12 data structure.
	EncryptionPassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl EncryptionPassphraseID,omitempty"`

	// The passphrase to use for integrity checking and decrypting the uploaded PKCS#12 data structure.
	Passphrase string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Passphrase,omitempty"`
}

type UploadCertificateWithPrivateKeyInPKCS12Response struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCertificateWithPrivateKeyInPKCS12Response"`

	// The certification path ID of the uploaded certification path.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`

	// The key ID of the uploaded key pair.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type GetKeyStatus struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetKeyStatus"`

	// The ID of the key for which to return the status.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type GetKeyStatusResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetKeyStatusResponse"`

	// Status of the requested key. The value should be one of the values in the tas:KeyStatus enumeration.
	KeyStatus string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyStatus,omitempty"`
}

type GetPrivateKeyStatus struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetPrivateKeyStatus"`

	// The ID of the key pair for which to return whether it contains a private key.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type GetPrivateKeyStatusResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetPrivateKeyStatusResponse"`

	// True if and only if the key pair contains a private key.
	HasPrivateKey bool `xml:"hasPrivateKey,omitempty"`
}

type GetAllKeys struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllKeys"`
}

type GetAllKeysResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllKeysResponse"`

	// Information about a key in the keystore.
	KeyAttribute KeyAttribute `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyAttribute,omitempty"`
}

type DeleteKey struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteKey"`

	// The ID of the key that is to be deleted from the keystore.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type DeleteKeyResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteKeyResponse"`
}

type CreatePKCS10CSR struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreatePKCS10CSR"`

	// The subject to be included in the CSR.
	Subject DistinguishedName `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Subject,omitempty"`

	// The ID of the key for which the CSR shall be created.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`

	// An attribute to be included in the CSR.
	CSRAttribute CSRAttribute `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CSRAttribute,omitempty"`

	// The signature algorithm to be used to sign the CSR.
	SignatureAlgorithm AlgorithmIdentifier `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SignatureAlgorithm,omitempty"`
}

type CreatePKCS10CSRResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreatePKCS10CSRResponse"`

	// The DER encoded PKCS#10 certification request.
	PKCS10CSR Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PKCS10CSR,omitempty"`
}

type CreateSelfSignedCertificate struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateSelfSignedCertificate"`

	// The X.509 version that the generated certificate shall comply to.
	X509Version PositiveInteger `xml:"http://www.onvif.org/ver10/pacs X509Version,omitempty"`

	// Distinguished name of the entity that the certificate shall belong to.
	Subject DistinguishedName `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Subject,omitempty"`

	// The ID of the key for which the certificate shall be created.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`

	// The client-defined alias of the certificate to be created.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// The X.509 not valid before information to be included in the certificate. Defaults to the device's current time or a time before the device's current time.
	NotValidBefore string `xml:"notValidBefore,omitempty"`

	// The X.509 not valid after information to be included in the certificate. Defaults to the time 99991231235959Z as specified in RFC 5280.
	NotValidAfter string `xml:"notValidAfter,omitempty"`

	// The signature algorithm to be used for signing the certificate.
	SignatureAlgorithm AlgorithmIdentifier `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SignatureAlgorithm,omitempty"`

	// An X.509v3 extension to be included in the certificate.
	Extension X509v3Extension `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Extension,omitempty"`
}

type CreateSelfSignedCertificateResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateSelfSignedCertificateResponse"`

	// The ID of the generated certificate.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`
}

type UploadCertificate struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCertificate"`

	// The base64-encoded DER representation of the X.509 certificate to be uploaded.
	Certificate Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Certificate,omitempty"`

	// The client-defined alias of the certificate.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// The client-defined alias of the key pair.
	KeyAlias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyAlias,omitempty"`

	// Indicates if the device shall verify that a matching key pair with a private key exists in the keystore.
	PrivateKeyRequired bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PrivateKeyRequired,omitempty"`
}

type UploadCertificateResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCertificateResponse"`

	// The ID of the uploaded certificate.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`

	// The ID of the key that the uploaded certificate certifies.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`
}

type GetCertificate struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertificate"`

	// The ID of the certificate to retrieve.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`
}

type GetCertificateResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertificateResponse"`

	// The DER representation of the certificate.
	Certificate X509Certificate `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Certificate,omitempty"`
}

type GetAllCertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertificates"`
}

type GetAllCertificatesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertificatesResponse"`

	// A certificate stored in the keystore.
	Certificate X509Certificate `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Certificate,omitempty"`
}

type DeleteCertificate struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertificate"`

	// The ID of the certificate to delete.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`
}

type DeleteCertificateResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertificateResponse"`
}

type CreateCertificationPath struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateCertificationPath"`

	// The IDs of the certificates to include in the certification path, where each certificate signature except for the last one in the path must be verifiable with the public key certified by the next certificate in the path.
	CertificateIDs CertificateIDs `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateIDs,omitempty"`

	// The client-defined alias of the certification path.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`
}

type CreateCertificationPathResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateCertificationPathResponse"`

	// The ID of the generated certification path.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type GetCertificationPath struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertificationPath"`

	// The ID of the certification path to retrieve.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type GetCertificationPathResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertificationPathResponse"`

	// The certification path that is stored under the given ID in the keystore.
	CertificationPath CertificationPath `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPath,omitempty"`
}

type GetAllCertificationPaths struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertificationPaths"`
}

type GetAllCertificationPathsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertificationPathsResponse"`

	// An ID of a certification path in the keystore.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type DeleteCertificationPath struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertificationPath"`

	// The ID of the certification path to delete.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type DeleteCertificationPathResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertificationPathResponse"`
}

type UploadPassphrase struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadPassphrase"`

	// The passphrase to upload.
	Passphrase string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Passphrase,omitempty"`

	// The alias for the passphrase to upload.
	PassphraseAlias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseAlias,omitempty"`
}

type UploadPassphraseResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadPassphraseResponse"`

	// The PassphraseID of the uploaded passphrase.
	PassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseID,omitempty"`
}

type GetAllPassphrases struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllPassphrases"`
}

type GetAllPassphrasesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllPassphrasesResponse"`

	// Information about a passphrase in the keystore.
	PassphraseAttribute PassphraseAttribute `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseAttribute,omitempty"`
}

type DeletePassphrase struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeletePassphrase"`

	// The ID of the passphrase that is to be deleted from the keystore.
	PassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseID,omitempty"`
}

type DeletePassphraseResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeletePassphraseResponse"`
}

type AddServerCertificateAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddServerCertificateAssignment"`

	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type AddServerCertificateAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddServerCertificateAssignmentResponse"`
}

type RemoveServerCertificateAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RemoveServerCertificateAssignment"`

	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type RemoveServerCertificateAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RemoveServerCertificateAssignmentResponse"`
}

type ReplaceServerCertificateAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl ReplaceServerCertificateAssignment"`

	OldCertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl OldCertificationPathID,omitempty"`

	NewCertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl NewCertificationPathID,omitempty"`
}

type ReplaceServerCertificateAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl ReplaceServerCertificateAssignmentResponse"`
}

type GetAssignedServerCertificates struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAssignedServerCertificates"`
}

type GetAssignedServerCertificatesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAssignedServerCertificatesResponse"`

	// The IDs of all certification paths that are assigned to the TLS server on the device.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`
}

type SetEnabledTLSVersions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetEnabledTLSVersions"`

	// List of TLS versions to allow.
	Versions TLSVersions `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Versions,omitempty"`
}

type SetEnabledTLSVersionsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetEnabledTLSVersionsResponse"`
}

type GetEnabledTLSVersions struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetEnabledTLSVersions"`
}

type GetEnabledTLSVersionsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetEnabledTLSVersionsResponse"`

	// List of allowed TLS versions.
	Versions TLSVersions `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Versions,omitempty"`
}

type UploadCRL struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCRL"`

	//
	// The CRL to be uploaded to the device.
	//
	Crl Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Crl,omitempty"`

	//
	// The alias to assign to the uploaded CRL.
	//
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	AnyParameters struct {
	} `xml:"anyParameters,omitempty"`
}

type UploadCRLResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UploadCRLResponse"`

	//
	// The ID of the uploaded CRL.
	//
	CrlID CRLID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CrlID,omitempty"`
}

type GetCRL struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCRL"`

	//
	// The ID of the CRL to be returned.
	//
	CrlID CRLID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CrlID,omitempty"`
}

type GetCRLResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCRLResponse"`

	//
	// The CRL with the requested ID.
	//
	Crl CRL `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Crl,omitempty"`
}

type GetAllCRLs struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCRLs"`
}

type GetAllCRLsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCRLsResponse"`

	//
	// A list of all CRLs that are stored in the keystore on the device.
	//
	Crl CRL `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Crl,omitempty"`
}

type DeleteCRL struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCRL"`

	//
	// The ID of the CRL to be deleted.
	//
	CrlID CRLID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CrlID,omitempty"`
}

type DeleteCRLResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCRLResponse"`
}

type CreateCertPathValidationPolicy struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateCertPathValidationPolicy"`

	//
	// The alias to assign to the created certification path validation policy.
	//
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	//
	// The parameters of the certification path validation policy to be created.
	//
	Parameters CertPathValidationParameters `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Parameters,omitempty"`

	//
	// The trust anchors of the certification path validation policy to be created.
	//
	TrustAnchor TrustAnchor `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl TrustAnchor,omitempty"`

	AnyParameters struct {
	} `xml:"anyParameters,omitempty"`
}

type CreateCertPathValidationPolicyResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CreateCertPathValidationPolicyResponse"`

	//
	// The ID of the created certification path validation policy.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type GetCertPathValidationPolicy struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertPathValidationPolicy"`

	//
	// The ID of the certification path validation policy to be created.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type GetCertPathValidationPolicyResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetCertPathValidationPolicyResponse"`

	//
	// The certification path validation policy that is stored under the requested ID.
	//
	CertPathValidationPolicy CertPathValidationPolicy `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicy,omitempty"`
}

type GetAllCertPathValidationPolicies struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertPathValidationPolicies"`
}

type GetAllCertPathValidationPoliciesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllCertPathValidationPoliciesResponse"`

	//
	// A list of all certification path validation policies that are stored in the keystore on the device.
	//
	CertPathValidationPolicy CertPathValidationPolicy `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicy,omitempty"`
}

type DeleteCertPathValidationPolicy struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertPathValidationPolicy"`

	//
	// The ID of the certification path validation policy to be deleted.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type DeleteCertPathValidationPolicyResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteCertPathValidationPolicyResponse"`
}

type SetClientAuthenticationRequired struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetClientAuthenticationRequired"`

	ClientAuthenticationRequired bool `xml:"clientAuthenticationRequired,omitempty"`
}

type SetClientAuthenticationRequiredResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetClientAuthenticationRequiredResponse"`
}

type GetClientAuthenticationRequired struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetClientAuthenticationRequired"`
}

type GetClientAuthenticationRequiredResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetClientAuthenticationRequiredResponse"`

	ClientAuthenticationRequired bool `xml:"clientAuthenticationRequired,omitempty"`
}

type AddCertPathValidationPolicyAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddCertPathValidationPolicyAssignment"`

	//
	// The ID of the certification path validation policy to assign to the TLS server.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type AddCertPathValidationPolicyAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddCertPathValidationPolicyAssignmentResponse"`
}

type RemoveCertPathValidationPolicyAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RemoveCertPathValidationPolicyAssignment"`

	//
	// The ID of the certification path validation policy to de-assign from the TLS server.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type RemoveCertPathValidationPolicyAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RemoveCertPathValidationPolicyAssignmentResponse"`
}

type ReplaceCertPathValidationPolicyAssignment struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl ReplaceCertPathValidationPolicyAssignment"`

	//
	// The ID of the certification path validation policy to be de-assigned from the TLS server.
	//
	OldCertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl OldCertPathValidationPolicyID,omitempty"`

	//
	// The ID of the certification path validation policy to assign to the TLS server.
	//
	NewCertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl NewCertPathValidationPolicyID,omitempty"`
}

type ReplaceCertPathValidationPolicyAssignmentResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl ReplaceCertPathValidationPolicyAssignmentResponse"`
}

type GetAssignedCertPathValidationPolicies struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAssignedCertPathValidationPolicies"`
}

type GetAssignedCertPathValidationPoliciesResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAssignedCertPathValidationPoliciesResponse"`

	//
	// A list of IDs of the certification path validation policies that are assigned to the TLS server.
	//
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`
}

type AddDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddDot1XConfiguration"`

	// The desired 802.1X configuration.
	Dot1XConfiguration Dot1XConfiguration `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XConfiguration,omitempty"`
}

type AddDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl AddDot1XConfigurationResponse"`

	// The unique identifier of the created 802.1X configuration.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`
}

type GetAllDot1XConfigurations struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllDot1XConfigurations"`
}

type GetAllDot1XConfigurationsResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetAllDot1XConfigurationsResponse"`

	// The list of unique identifiers of 802.1X configurations on the device.
	Configuration Dot1XConfiguration `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Configuration,omitempty"`
}

type GetDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetDot1XConfiguration"`

	// The unique identifier of the desired 802.1X configuration.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`
}

type GetDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetDot1XConfigurationResponse"`

	// The 802.1X configuration, without password information.
	Dot1XConfiguration Dot1XConfiguration `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XConfiguration,omitempty"`
}

type DeleteDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteDot1XConfiguration"`

	// The unique identifier of the 802.1X configuration to be deleted.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`
}

type DeleteDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteDot1XConfigurationResponse"`
}

type SetNetworkInterfaceDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetNetworkInterfaceDot1XConfiguration"`

	// The unique identifier of the Network Interface on which the 802.1X configuration is to be set. (NOTE: the network interface token is defined in devicemgmt.wsdl as tt:ReferenceToken, which is a derived type of xs:string.  To avoid importing all of common.xsd for this single type, the base type is used here.)
	Token string `xml:"token,omitempty"`

	// The unique identifier of the 802.1X configuration to be set.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`
}

type SetNetworkInterfaceDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SetNetworkInterfaceDot1XConfigurationResponse"`

	// Indicates whether or not a reboot is required after configuration updates.
	RebootNeeded bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RebootNeeded,omitempty"`
}

type GetNetworkInterfaceDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetNetworkInterfaceDot1XConfiguration"`

	// The unique identifier of the Network Interface for which the 802.1X configuration is to be retrieved. (NOTE: the network interface token is defined in devicemgmt.wsdl as tt:ReferenceToken, which is a derived type of xs:string.  To avoid importing all of common.xsd for this single type, the base type is used here.)
	Token string `xml:"token,omitempty"`
}

type GetNetworkInterfaceDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GetNetworkInterfaceDot1XConfigurationResponse"`

	// The unique identifier of 802.1X configuration assigned to the Network Interface.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`
}

type DeleteNetworkInterfaceDot1XConfiguration struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteNetworkInterfaceDot1XConfiguration"`

	// The unique identifier of the Network Interface for which the 802.1X configuration is to be deleted. (NOTE: the network interface token is defined in devicemgmt.wsdl as tt:ReferenceToken, which is a derived type of xs:string.  To avoid importing all of common.xsd for this single type, the base type is used here.)
	Token string `xml:"token,omitempty"`
}

type DeleteNetworkInterfaceDot1XConfigurationResponse struct {
	XMLName xml.Name `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DeleteNetworkInterfaceDot1XConfigurationResponse"`

	// Indicates whether or not a reboot is required after configuration updates.
	RebootNeeded bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RebootNeeded,omitempty"`
}

type KeyAttribute struct {

	// The ID of the key.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`

	// The client-defined alias of the key.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// Absent if the key is not a key pair. True if and only if the key is a key pair and contains a private key. False if and only if the key is a key pair and does not contain a private key.
	HasPrivateKey bool `xml:"hasPrivateKey,omitempty"`

	// The status of the key. The value should be one of the values in the tas:KeyStatus enumeration.
	KeyStatus string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyStatus,omitempty"`

	// True if and only if the key was generated outside the device.
	ExternallyGenerated bool `xml:"externallyGenerated,omitempty"`

	// True if and only if the key is stored in a specially protected hardware component inside the device.
	SecurelyStored bool `xml:"securelyStored,omitempty"`

	Extension struct {
	} `xml:"Extension,omitempty"`
}

type DNAttributeTypeAndValue struct {

	// The attribute type.
	Type DNAttributeType `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Type,omitempty"`

	// The value of the attribute.
	Value DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Value,omitempty"`
}

type MultiValuedRDN struct {

	// A list of types and values defining a multi-valued RDN
	Attribute DNAttributeTypeAndValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Attribute,omitempty"`
}

type DistinguishedName struct {

	// A country name as specified in
	// X.500.
	Country DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Country,omitempty"`

	// An organization name as specified in
	// X.500.
	Organization DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Organization,omitempty"`

	// An organizational unit name as specified in
	// X.500.
	OrganizationalUnit DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl OrganizationalUnit,omitempty"`

	// A distinguished name qualifier as specified in
	// X.500.
	DistinguishedNameQualifier DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DistinguishedNameQualifier,omitempty"`

	// A state or province name as specified in
	// X.500.
	StateOrProvinceName DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl StateOrProvinceName,omitempty"`

	// A common name as specified in
	// X.500.
	CommonName DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CommonName,omitempty"`

	// A serial number as specified in
	// X.500.
	SerialNumber DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SerialNumber,omitempty"`

	// A locality as specified in X.500.
	Locality DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Locality,omitempty"`

	// A title as specified in X.500.
	Title DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Title,omitempty"`

	// A surname as specified in X.500.
	Surname DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Surname,omitempty"`

	// A given name as specified in X.500.
	GivenName DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GivenName,omitempty"`

	// Initials as specified in X.500.
	Initials DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Initials,omitempty"`

	// A pseudonym as specified in X.500.
	Pseudonym DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Pseudonym,omitempty"`

	// A generation qualifier as specified in
	// X.500.
	GenerationQualifier DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GenerationQualifier,omitempty"`

	// A generic type-value pair
	// attribute.
	GenericAttribute DNAttributeTypeAndValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl GenericAttribute,omitempty"`

	// A multi-valued RDN
	MultiValuedRDN MultiValuedRDN `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl MultiValuedRDN,omitempty"`

	AnyAttribute struct {

		// Domain Component as specified in RFC3739
		DomainComponent DNAttributeValue `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl DomainComponent,omitempty"`
	} `xml:"anyAttribute,omitempty"`
}

type AlgorithmIdentifier struct {

	// The OID of the algorithm in dot-decimal form.
	Algorithm DotDecimalOID `xml:"algorithm,omitempty"`

	// Optional parameters of the algorithm (depending on the algorithm).
	Parameters Base64DERencodedASN1Value `xml:"parameters,omitempty"`

	AnyParameters struct {
	} `xml:"anyParameters,omitempty"`
}

type BasicRequestAttribute struct {

	// The OID of the attribute.
	OID DotDecimalOID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl OID,omitempty"`

	// The value of the attribute as a base64-encoded DER representation of an ASN.1 value.
	Value Base64DERencodedASN1Value `xml:"value,omitempty"`
}

type CSRAttribute struct {

	// An X.509v3 extension field.
	X509v3Extension X509v3Extension `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl X509v3Extension,omitempty"`

	// A basic CSR attribute.
	BasicRequestAttribute BasicRequestAttribute `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl BasicRequestAttribute,omitempty"`

	AnyAttribute struct {
	} `xml:"anyAttribute,omitempty"`
}

type X509v3Extension struct {

	// The OID of the extension field.
	ExtnOID DotDecimalOID `xml:"extnOID,omitempty"`

	// True if and only if the extension is critical.
	Critical bool `xml:"critical,omitempty"`

	// The value of the extension field as a base64-encoded DER representation of an ASN.1 value.
	ExtnValue Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl extnValue,omitempty"`
}

type X509Certificate struct {

	// The ID of the certificate.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`

	// The ID of the key that this certificate associates to the certificate subject.
	KeyID KeyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeyID,omitempty"`

	// The client-defined alias of the certificate.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// The base64-encoded DER representation of the X.509 certificate.
	CertificateContent Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateContent,omitempty"`
}

type CertificateIDs struct {

	// A certificate ID.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`
}

type CertificationPath struct {

	// A certificate in the certification path.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`

	// The client-defined alias of the certification path.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	AnyElement struct {
	} `xml:"anyElement,omitempty"`
}

type PassphraseAttribute struct {

	// The ID of the passphrase.
	PassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseID,omitempty"`

	// The alias of the passphrase.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`
}

type Dot1XCapabilities struct {

	// The maximum number of 802.1X configurations that may be defined simultaneously.

	MaximumNumberOfDot1XConfigurations PositiveInteger `xml:"http://www.onvif.org/ver10/pacs MaximumNumberOfDot1XConfigurations,attr,omitempty"`

	// The authentication methods supported by the 802.1X implementation.

	Dot1XMethods Dot1XMethods `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XMethods,attr,omitempty"`
}

type Dot1XStage struct {

	// The identity used in this authentication method, if required.
	Identity string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Identity,omitempty"`

	// The unique identifier of the certification path used in this authentication method, if required.
	CertificationPathID CertificationPathID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificationPathID,omitempty"`

	// The identifier for the password used in this authentication method, if required.  If Identity is used as an anonymous identity for this authentication method, PassphraseID is ignored.
	PassphraseID PassphraseID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PassphraseID,omitempty"`

	// The configuration of the next stage of authentication, if required.
	Inner *Dot1XStage `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Inner,omitempty"`

	Extension Dot1XStageExtension `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Extension,omitempty"`

	// The authentication method for this stage (e.g., "EAP-PEAP").

	Method string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Method,attr,omitempty"`
}

type Dot1XStageExtension struct {
}

type Dot1XConfiguration struct {

	// The unique identifier of the IEEE 802.1X configuration.
	Dot1XID Dot1XID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XID,omitempty"`

	// The client-defined alias of the 802.1X configuration.
	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	// The outer level authentication method used in this 802.1X configuration.
	Outer Dot1XStage `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Outer,omitempty"`
}

type CRL struct {
	CRLID CRLID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CRLID,omitempty"`

	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	CRLContent Base64DERencodedASN1Value `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CRLContent,omitempty"`
}

type CertPathValidationParameters struct {

	// True if and only if the TLS server shall not authenticate client certificates that do not contain the TLS WWW client authentication key usage extension as specified in RFC 5280, Sect. 4.2.1.12.
	RequireTLSWWWClientAuthExtendedKeyUsage bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RequireTLSWWWClientAuthExtendedKeyUsage,omitempty"`

	// True if and only if delta CRLs, if available, shall be applied to CRLs.
	UseDeltaCRLs bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl UseDeltaCRLs,omitempty"`

	AnyParameters struct {
	} `xml:"anyParameters,omitempty"`
}

type TrustAnchor struct {

	// The certificate ID of the certificate to be used as trust anchor.
	CertificateID CertificateID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertificateID,omitempty"`
}

type CertPathValidationPolicy struct {
	CertPathValidationPolicyID CertPathValidationPolicyID `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl CertPathValidationPolicyID,omitempty"`

	Alias string `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Alias,omitempty"`

	Parameters CertPathValidationParameters `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Parameters,omitempty"`

	TrustAnchor TrustAnchor `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl TrustAnchor,omitempty"`

	AnyParameters struct {
	} `xml:"anyParameters,omitempty"`
}

type KeystoreCapabilities struct {

	// The signature algorithms supported by the keystore implementation.
	SignatureAlgorithms AlgorithmIdentifier `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SignatureAlgorithms,omitempty"`

	AnyElement struct {
	} `xml:"anyElement,omitempty"`

	// Indicates the maximum number of keys that the device can store simultaneously.

	MaximumNumberOfKeys PositiveInteger `xml:"http://www.onvif.org/ver10/pacs MaximumNumberOfKeys,attr,omitempty"`

	// Indicates the maximum number of certificates that the device can store simultaneously.

	MaximumNumberOfCertificates PositiveInteger `xml:"http://www.onvif.org/ver10/pacs MaximumNumberOfCertificates,attr,omitempty"`

	// Indicates the maximum number of certification paths that the device can store simultaneously.

	MaximumNumberOfCertificationPaths PositiveInteger `xml:"http://www.onvif.org/ver10/pacs MaximumNumberOfCertificationPaths,attr,omitempty"`

	// Indication that the device supports on-board RSA key pair generation.

	RSAKeyPairGeneration bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RSAKeyPairGeneration,attr,omitempty"`

	// Indicates which RSA key lengths are supported by the device.

	RSAKeyLengths RSAKeyLengths `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl RSAKeyLengths,attr,omitempty"`

	// Indicates support for creating PKCS#10 requests for RSA keys and uploading the certificate obtained from a CA..

	PKCS10ExternalCertificationWithRSA bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PKCS10ExternalCertificationWithRSA,attr,omitempty"`

	// Indicates support for creating self-signed certificates for RSA keys.

	SelfSignedCertificateCreationWithRSA bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl SelfSignedCertificateCreationWithRSA,attr,omitempty"`

	// Indicates which X.509 versions are supported by the device.

	X509Versions X509Versions `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl X509Versions,attr,omitempty"`

	// Indicates the maximum number of passphrases that the device is able to store simultaneously.

	MaximumNumberOfPassphrases NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfPassphrases,attr,omitempty"`

	// Indicates support for uploading an RSA key pair in a PKCS#8 data structure.

	PKCS8RSAKeyPairUpload bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PKCS8RSAKeyPairUpload,attr,omitempty"`

	// Indicates support for uploading a certificate along with an RSA private key in a PKCS#12 data structure.

	PKCS12CertificateWithRSAPrivateKeyUpload bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PKCS12CertificateWithRSAPrivateKeyUpload,attr,omitempty"`

	// Indicates which password-based encryption algorithms are supported by the device.

	PasswordBasedEncryptionAlgorithms PasswordBasedEncryptionAlgorithms `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PasswordBasedEncryptionAlgorithms,attr,omitempty"`

	// Indicates which password-based MAC algorithms are supported by the device.

	PasswordBasedMACAlgorithms PasswordBasedMACAlgorithms `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl PasswordBasedMACAlgorithms,attr,omitempty"`

	// Indicates the maximum number of CRLs that the device is able to store simultaneously.

	MaximumNumberOfCRLs NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfCRLs,attr,omitempty"`

	// Indicates the maximum number of certification path validation policies that the device is able to store simultaneously.

	MaximumNumberOfCertificationPathValidationPolicies NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfCertificationPathValidationPolicies,attr,omitempty"`

	// Indicates whether a device supports checking for the TLS WWW client auth extended key usage extension while validating certification paths.

	EnforceTLSWebClientAuthExtKeyUsage bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl EnforceTLSWebClientAuthExtKeyUsage,attr,omitempty"`

	// Indicates the device requires that each certificate with private key has its own unique key.

	NoPrivateKeySharing bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl NoPrivateKeySharing,attr,omitempty"`
}

type TLSServerCapabilities struct {

	// Indicates which TLS versions are supported by the device.

	TLSServerSupported TLSVersions `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl TLSServerSupported,attr,omitempty"`

	// Indicates whether the device supports enabling and disabling specific TLS versions.

	EnabledVersionsSupported bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl EnabledVersionsSupported,attr,omitempty"`

	// Indicates the maximum number of certification paths that may be assigned to the TLS server simultaneously.

	MaximumNumberOfTLSCertificationPaths PositiveInteger `xml:"http://www.onvif.org/ver10/pacs MaximumNumberOfTLSCertificationPaths,attr,omitempty"`

	// Indicates whether the device supports TLS client authentication.

	TLSClientAuthSupported bool `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl TLSClientAuthSupported,attr,omitempty"`

	// Indicates the maximum number of certification path validation policies that may be assigned to the TLS server simultaneously.

	MaximumNumberOfTLSCertificationPathValidationPolicies NonNegativeInteger `xml:"http://www.onvif.org/ver10/schema MaximumNumberOfTLSCertificationPathValidationPolicies,attr,omitempty"`
}

type Capabilities struct {

	// The capabilities of the keystore implementation.
	KeystoreCapabilities KeystoreCapabilities `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl KeystoreCapabilities,omitempty"`

	// The capabilities of the TLS server implementation.
	TLSServerCapabilities TLSServerCapabilities `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl TLSServerCapabilities,omitempty"`

	// The capabilities of the 802.1X implementation.
	Dot1XCapabilities Dot1XCapabilities `xml:"http://www.onvif.org/ver10/advancedsecurity/wsdl Dot1XCapabilities,omitempty"`
}

type AdvancedSecurityService interface {

	/* Returns the capabilities of the security configuraiton service. The result is returned in a typed answer. */
	GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)

	GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error)
}

type advancedSecurityService struct {
	client *soap.Client
	xaddr  string
}

func NewAdvancedSecurityService(client *soap.Client, xaddr string) AdvancedSecurityService {
	return &advancedSecurityService{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *advancedSecurityService) GetServiceCapabilitiesContext(ctx context.Context, request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	response := new(GetServiceCapabilitiesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetServiceCapabilities", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *advancedSecurityService) GetServiceCapabilities(request *GetServiceCapabilities) (*GetServiceCapabilitiesResponse, error) {
	return service.GetServiceCapabilitiesContext(
		context.Background(),
		request,
	)
}

type Keystore interface {

	/*
		This operation triggers the asynchronous generation of an RSA key pair of a particular key length (specified as the number of bits) as specified in [RFC 3447], with a suitable key generation mechanism on the device.
		Keys, especially RSA key pairs, are uniquely identified using key IDs.
		If the device does not have not enough storage capacity for storing the key pair to be created, the maximum number of keys reached fault shall be produced and no key pair shall be generated.
		Otherwise, the operation generates a keyID for the new key and associates the generating status to it.
		Immediately after key generation has started, the device shall return the keyID to the client and continue to generate the key pair.
		The client may query the device with the GetKeyStatus operation whether the generation has finished.
		The client may also subscribe to Key Status events to be notified about key status changes.
		The device also returns a best-effort estimate of how much time it requires to create the key pair.
		A client may use this information as an indication how long to wait before querying the device whether key generation is completed.
		After the key has been successfully created, the device shall assign it the ok status. If the key generation fails, the device shall assign the key the corrupt status.
	*/
	CreateRSAKeyPair(request *CreateRSAKeyPair) (*CreateRSAKeyPairResponse, error)

	CreateRSAKeyPairContext(ctx context.Context, request *CreateRSAKeyPair) (*CreateRSAKeyPairResponse, error)

	/*
	   This operation uploads a key pair in a PKCS#8 data structure as specified in [RFC 5958, RFC 5959].
	   If an encryption passphrase ID is supplied in the request, the device shall assume that the KeyPair parameter contains an EncryptedPrivateKeyInfo ASN.1
	   structure that is encrypted under the passphrase in the keystore that corresponds to the supplied ID, where the EncryptedPrivateKeyInfo structure contains
	   both the private key and the corresponding public key. If no encryption passphrase ID is supplied, the device shall assume that the KeyPair parameter contains a
	   OneAsymmetricKey ASN.1 structure which contains both the private key and the corresponding public key.
	*/
	UploadKeyPairInPKCS8(request *UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, error)

	UploadKeyPairInPKCS8Context(ctx context.Context, request *UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, error)

	/*
	   This operation uploads a certification path consisting of X.509 certificates as specified by [RFC 5280] in DER encoding along with a private key to a device’s keystore.
	   Certificates and private key are supplied in the form of a PKCS#12 file as specified in [PKCS#12].
	   The device shall support PKCS#12 files that contain the following safe bags:
	   *	one or more instances of CertBag [PKCS#12, Sect. 4.2.3]
	   *	either exactly one instance of KeyBag [PKCS#12, Sect. 4.3.1] or exactly one instance of PKCS8ShroudedKeyBag [PKCS#12, Sect. 4.2.2].
	   If the IgnoreAdditionalCertificates parameter has the value true, the device shall behave as if the client had supplied only the first CertBag in the sequence of CertBag instances.
	   The device shall support PKCS#12 passphrase integrity mode for integrity protection of the PKCS#12 PFX as specified in [PKCS#12, Sect. 4].
	   The device shall support PKCS8ShroudedKeyBags that are encrypted with the same passphrase as the CertBag instances.
	   If an integrity passphrase ID is supplied, the device shall use the corresponding passphrase in the keystore to check the integrity of the supplied PKCS#12 PFX.
	   If an integrity passphrase ID is supplied, but the supplied PKCS#12 PFX has no integrity protection, the device shall produce a BadPKCS12File fault and shall
	   not store the uploaded certificates nor the uploaded key pair in the keystore.
	   If an encryption passphrase ID is supplied, the device shall use the corresponding passphrase in the keystore to decrypt the PKCS8ShroudedKeyBag and the CertBag instances.
	   If an EncryptionPassphraseID is supplied, but a CertBag is not encrypted, the device shall ignore the supplied EncryptionPassphraseID when processing this CertBag.
	   If an EncryptionPassphraseID is supplied, but a KeyBag is provided instead of a PKCS8ShroudedKeyBag, the device shall ignore the supplied EncryptionPassphraseID when processing the KeyBag.
	*/
	UploadCertificateWithPrivateKeyInPKCS12(request *UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, error)

	UploadCertificateWithPrivateKeyInPKCS12Context(ctx context.Context, request *UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, error)

	/*
		This operation returns the status of a key.
		Keys are uniquely identified using key IDs. If no key is stored under the requested key ID in the keystore, an InvalidKeyID fault is produced.
		Otherwise, the status of the key is returned.
	*/
	GetKeyStatus(request *GetKeyStatus) (*GetKeyStatusResponse, error)

	GetKeyStatusContext(ctx context.Context, request *GetKeyStatus) (*GetKeyStatusResponse, error)

	/*
		This operation returns whether a key pair contains a private key.
		Keys are uniquely identified using key IDs. If no key is stored under the requested key ID in the keystore or the key identified by the requested key ID does not identify a key pair,
		the device shall produce an InvalidKeyID fault.
		Otherwise, this operation returns true if the key pair identified by the key ID contains a private key, and false otherwise.
	*/
	GetPrivateKeyStatus(request *GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, error)

	GetPrivateKeyStatusContext(ctx context.Context, request *GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, error)

	/*
		This operation returns information about all keys that are stored in the device’s keystore.
		This operation may be used, e.g., if a client lost track of which keys are present on the device.
		If no key is stored on the device, an empty list is returned.
	*/
	GetAllKeys(request *GetAllKeys) (*GetAllKeysResponse, error)

	GetAllKeysContext(ctx context.Context, request *GetAllKeys) (*GetAllKeysResponse, error)

	/*
		This operation deletes a key from the device’s keystore.
		Keys are uniquely identified using key IDs. If no key is stored under the requested key ID in the keystore, a device shall produce an InvalidArgVal fault.
		If a reference exists for the specified key, a device shall produce the corresponding fault and shall not delete the key.
		If there is a key under the requested key ID stored in the keystore and the key could not be deleted, a device shall produce a KeyDeletion fault.
		If the key has the status generating, a device shall abort the generation of the key and delete from the keystore all data generated for this key.
		After a key is successfully deleted, the device may assign its former ID to other keys.
	*/
	DeleteKey(request *DeleteKey) (*DeleteKeyResponse, error)

	DeleteKeyContext(ctx context.Context, request *DeleteKey) (*DeleteKeyResponse, error)

	/*
		This operation generates a DER-encoded PKCS#10 v1.7 certification request (sometimes also called certificate signing request or CSR) as specified in RFC 2986
		for a public key on the device.
		The key pair that contains the public key for which a certification request shall be produced is specified by its key ID.
		If no key is stored under the requested KeyID or the key specified by the requested KeyID is not an asymmetric key pair, an invalid key ID fault shall be produced and
		no CSR shall be generated.

		A device that supports this command shall as minimum support the sha-1WithRSAEncryption signature algorithm as specified in RFC 3279.
		If the specified signature algorithm is not supported by the device, an UnsupportedSignatureAlgorithm fault shall be produced and no CSR shall be generated.

		If the public key identified by the requested Key ID is an invalid input to the specified signature algorithm, a KeySignatureAlgorithmMismatch fault shall be produced
		and no CSR shall be generated.
		If the key pair does not have status ok, a device shall produce an InvalidKeyStatus fault and no CSR shall be generated.
	*/
	CreatePKCS10CSR(request *CreatePKCS10CSR) (*CreatePKCS10CSRResponse, error)

	CreatePKCS10CSRContext(ctx context.Context, request *CreatePKCS10CSR) (*CreatePKCS10CSRResponse, error)

	/*
		This operation generates for a public key on the device a self-signed X.509 certificate that complies to RFC 5280.
		The X509Version parameter specifies the version of X.509 that the generated certificate shall comply to.
		A device that supports this command shall support the generation of X.509v3 certificates as specified in RFC 5280 and may additionally be able to handle other X.509 certificate formats
		as indicated by the X.509Versions capability.
		The key pair that contains the public key for which a self-signed certificate shall be produced is specified by its key pair ID.
		The subject parameter describes the entity that the public key belongs to.
		If the key pair does not have status ok, a device shall produce an InvalidKeyStatus fault and no certificate shall be generated.

		The signature algorithm parameter determines which signature algorithm shall be used for signing the certification request with the public key specified by the key ID parameter.
		A device that supports this command shall as minimum support the sha-1WithRSAEncryption signature algorithm as specified in RFC 3279.
		The Extensions parameter specifies potential X509v3 extensions that shall be contained in the certificate.
		A device that supports this command shall support the extensions that are defined in [RFC 5280], Sect. 4.2] as mandatory for CAs that issue self-signed certificates.

		Certificates are uniquely identified using certificate IDs. If the command was successful, the device generates a new ID for the generated certificate and returns this ID.
		If the device does not have not enough storage capacity for storing the certificate to be created, the maximum number of certificates reached fault shall be produced and no certificate shall be generated.
	*/
	CreateSelfSignedCertificate(request *CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, error)

	CreateSelfSignedCertificateContext(ctx context.Context, request *CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, error)

	/*
		This operation uploads an X.509 certificate as specified by [RFC 5280] in DER encoding and the public key in the certificate to a device’s keystore.
		A device that supports this command shall be able to handle X.509v3 certificates as specified in RFC 5280 and may additionally be able to handle other X.509 certificate formats as indicated by the X.509Versions capability.
		A device that supports this command shall support sha1-WithRSAEncryption as certificate signature algorithm.

		Certificates are uniquely identified using certificate IDs, and key pairs are uniquely identified using key IDs.
		The device shall generate a new certificate ID for the uploaded certificate.
		Certain certificate usages, e.g. TLS server authentication, require the private key that corresponds to the public key in the certificate to be present in the keystore.
		In such cases, the client may indicate that it expects the device to produce a fault if the matching private key for
		the uploaded certificate is not present in the keystore by setting the PrivateKeyRequired argument in the upload request to true.

		The uploaded certificate has to be linked to a key pair in the keystore.
		If no private key is required for the public key in the certificate and a key pair exists in the keystore with a public key equal to the public key in the certificate,
		the uploaded certificate is linked to the key pair identified by the supplied key ID by adding a reference from the certificate to the key pair.
		If no private key is required for the public key in the certificate and no key pair exists with the public key equal to the public key in the certificate,
		a new key pair with status ok is created with the public key from the certificate, and this key pair is linked to the uploaded certificate by adding a reference from
		the certificate to the key pair.
		If a private key is required for the public key in the certificate, and a key pair exists in the keystore with a private key that matches the public key in the certificate,
		the uploaded certificate is linked to this keypair by adding a reference from the certificate to the key pair.
		If a private key is required for the public key and no such keypair exists in the keystore, the NoMatchingPrivateKey fault shall be produced and the certificate
		shall not be stored in the keystore.
		If the key pair that the certificate shall be linked to does not have status ok, an InvalidKeyID fault is produced, and the uploaded certificate is not stored in the keystore.
		If the device cannot process the uploaded certificate, a BadCertificate fault is produced and neither the uploaded certificate nor the public key are stored in the device’s keystore.
		The BadCertificate fault shall not be produced based on the mere fact that the device’s current time lies outside the interval defined by the notBefore and notAfter fields as specified by [RFC 5280], Sect. 4.1 .
		This operation shall not mark the uploaded certificate as trusted.

		If the device does not have not enough storage capacity for storing the certificate to be uploaded, the maximum number of certificates reached fault shall be produced
		and no certificate shall be uploaded.
		If the device does not have not enough storage capacity for storing the key pair that eventually has to be created, the device shall generate a maximum number of keys reached fault.
		Furthermore the device shall not generate a key pair and no certificate shall be stored.
	*/
	UploadCertificate(request *UploadCertificate) (*UploadCertificateResponse, error)

	UploadCertificateContext(ctx context.Context, request *UploadCertificate) (*UploadCertificateResponse, error)

	/*
		This operation returns a specific certificate from the device’s keystore.
		Certificates are uniquely identified using certificate IDs. If no certificate is stored under the requested certificate ID in the keystore, an InvalidArgVal fault is produced.
		It shall be noted that this command does not return the private key that is associated to the public key in the certificate.
	*/
	GetCertificate(request *GetCertificate) (*GetCertificateResponse, error)

	GetCertificateContext(ctx context.Context, request *GetCertificate) (*GetCertificateResponse, error)

	/*
		This operation returns the IDs of all certificates that are stored in the device’s keystore.
		This operation may be used, e.g.,  if a client lost track of which certificates are present on the device.
		If no certificate is stored in the device’s keystore, an empty list is returned.
	*/
	GetAllCertificates(request *GetAllCertificates) (*GetAllCertificatesResponse, error)

	GetAllCertificatesContext(ctx context.Context, request *GetAllCertificates) (*GetAllCertificatesResponse, error)

	/*
		This operation deletes a certificate from the device’s keystore.
		The operation shall not delete the public key that is contained in the certificate from the keystore.
		Certificates are uniquely identified using certificate IDs. If no certificate is stored under the requested certificate ID in the keystore, an InvalidArgVal fault is produced.
		If there is a certificate under the requested certificate ID stored in the keystore and the certificate could not be deleted, a CertificateDeletion fault is produced.
		If a reference exists for the specified certificate, the certificate shall not be deleted and the corresponding fault shall be produced.
		After a certificate has been  successfully deleted, the device may assign its former ID to other certificates.
	*/
	DeleteCertificate(request *DeleteCertificate) (*DeleteCertificateResponse, error)

	DeleteCertificateContext(ctx context.Context, request *DeleteCertificate) (*DeleteCertificateResponse, error)

	/*
		This operation creates a sequence of certificates that may be used, e.g., for certification path validation or for TLS server authentication.
		Certification paths are uniquely identified using certification path IDs. Certificates are uniquely identified using certificate IDs.
		A certification path contains a sequence of certificate IDs.
		If there is a certificate ID in the sequence of supplied certificate IDs for which no certificate exists in the device’s keystore, the corresponding fault shall be produced
		and no certification path shall be created.

		The signature of each certificate in the certification path except for the last one must be verifiable with the public key contained in the next certificate in the path.
		If there is a certificate ID in the request other than the last ID for which the corresponding certificate cannot be verified with the public key in the certificate identified
		by the next certificate ID, an InvalidCertificateChain fault shall be produced and no certification path shall be created.
	*/
	CreateCertificationPath(request *CreateCertificationPath) (*CreateCertificationPathResponse, error)

	CreateCertificationPathContext(ctx context.Context, request *CreateCertificationPath) (*CreateCertificationPathResponse, error)

	/*
		This operation returns a specific certification path from the device’s keystore.
		Certification paths are uniquely identified using certification path IDs.
		If no certification path is stored under the requested ID in the keystore, an InvalidArgVal fault is produced.
	*/
	GetCertificationPath(request *GetCertificationPath) (*GetCertificationPathResponse, error)

	GetCertificationPathContext(ctx context.Context, request *GetCertificationPath) (*GetCertificationPathResponse, error)

	/*
		This operation returns the IDs of all certification paths that are stored in the device’s keystore.
		This operation may be used, e.g., if a client lost track of which certificates are present on the device.
		If no certification path is stored on the device, an empty list is returned.
	*/
	GetAllCertificationPaths(request *GetAllCertificationPaths) (*GetAllCertificationPathsResponse, error)

	GetAllCertificationPathsContext(ctx context.Context, request *GetAllCertificationPaths) (*GetAllCertificationPathsResponse, error)

	/*
		This operation deletes a certification path from the device’s keystore.
		This operation shall not delete the certificates that are referenced by the certification path.
		Certification paths are uniquely identified using certification path IDs.
		If no certification path is stored under the requested certification path ID in the keystore, an InvalidArgVal fault is produced.
		If there is a certification path under the requested certification path ID stored in the keystore and the certification path could not be deleted,
		a CertificationPathDeletion fault is produced.
		If a reference exists for the specified certification path, the certification path shall not be deleted and the corresponding fault shall be produced.
		After a certification path is successfully deleted, the device may assign its former ID to other certification paths.
	*/
	DeleteCertificationPath(request *DeleteCertificationPath) (*DeleteCertificationPathResponse, error)

	DeleteCertificationPathContext(ctx context.Context, request *DeleteCertificationPath) (*DeleteCertificationPathResponse, error)

	/*
	   This operation uploads a passphrase to the keystore of the device.
	*/
	UploadPassphrase(request *UploadPassphrase) (*UploadPassphraseResponse, error)

	UploadPassphraseContext(ctx context.Context, request *UploadPassphrase) (*UploadPassphraseResponse, error)

	/*
	   This operation returns information about all passphrases that are stored in the keystore of the device.
	   This operation may be used, e.g., if a client lost track of which passphrases are present on the device.
	   If no passphrase is stored on the device, the device shall return an empty list.
	*/
	GetAllPassphrases(request *GetAllPassphrases) (*GetAllPassphrasesResponse, error)

	GetAllPassphrasesContext(ctx context.Context, request *GetAllPassphrases) (*GetAllPassphrasesResponse, error)

	/*
	   This operation deletes a passphrase from the keystore of the device.
	*/
	DeletePassphrase(request *DeletePassphrase) (*DeletePassphraseResponse, error)

	DeletePassphraseContext(ctx context.Context, request *DeletePassphrase) (*DeletePassphraseResponse, error)

	/*
	   This operation uploads a certificate revocation list (CRL) as specified in [RFC 5280] to the keystore on the device.
	   If the device does not have enough storage space to store the CRL to be uploaded, the device shall produce a MaximumNumberOfCRLsReached fault and shall not store the supplied CRL.
	   If the device is not able to process the supplied CRL, the device shall produce a BadCRL fault and shall not store the supplied CRL.
	   If the device does not support the signature algorithm that was used to sign the supplied CRL, the device shall produce an UnsupportedSignatureAlgorithm fault and shall not store the supplied CRL.
	*/
	UploadCRL(request *UploadCRL) (*UploadCRLResponse, error)

	UploadCRLContext(ctx context.Context, request *UploadCRL) (*UploadCRLResponse, error)

	/*
	   This operation returns a specific certificate revocation list (CRL) from the keystore on the device.
	   Certification revocation lists are uniquely identified using CRLIDs. If no CRL is stored under the requested CRLID, the device shall produce a CRLID fault.
	*/
	GetCRL(request *GetCRL) (*GetCRLResponse, error)

	GetCRLContext(ctx context.Context, request *GetCRL) (*GetCRLResponse, error)

	/*
	   This operation returns all certificate revocation lists (CRLs) that are stored in the keystore on the device.
	   If no certificate revocation list is stored in the device’s keystore, an empty list is returned.
	*/
	GetAllCRLs(request *GetAllCRLs) (*GetAllCRLsResponse, error)

	GetAllCRLsContext(ctx context.Context, request *GetAllCRLs) (*GetAllCRLsResponse, error)

	/*
	   This operation deletes a certificate revocation list (CRL) from the keystore on the device.
	   Certification revocation lists are uniquely identified using CRLIDs. If no CRL is stored under the requested CRLID, the device shall produce a CRLID fault.
	   If a reference exists for the specified CRL, the device shall produce a ReferenceExists fault and shall not delete the CRL.
	   After a CRL has been successfully deleted, a device may assign its former ID to other CRLs.
	*/
	DeleteCRL(request *DeleteCRL) (*DeleteCRLResponse, error)

	DeleteCRLContext(ctx context.Context, request *DeleteCRL) (*DeleteCRLResponse, error)

	/*
	   This operation creates a certification path validation policy.
	   Certification path validation policies are uniquely identified using certification path validation policy IDs. The device shall generate a new certification path validation policy ID for the created certification path validation policy.
	   For the certification path validation parameters that are not represented in the certPathValidationParameters data type, the device shall use the default values specified in Sect. 3.
	   If the device does not have enough storage capacity for storing the certification path validation policy to be created, the device shall produce a maximum number of certification path validation policies reached fault and shall not create a certification path validation policy.
	   If there is at least one trust anchor certificate ID in the request for which there exists no certificate in the device’s keystore, the device shall produce a CertificateID fault and shall not create a certification path validation policy.
	   If the device cannot process the supplied certification path validation parameters, the device shall produce a CertPathValidationParameters fault and shall not create a certification path validation policy.
	*/
	CreateCertPathValidationPolicy(request *CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, error)

	CreateCertPathValidationPolicyContext(ctx context.Context, request *CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, error)

	/*
	   This operation returns a certification path validation policy from the keystore on the device.
	   Certification path validation policies are uniquely identified using certification path validation policy IDs. If no certification path validation policy is stored under the requested certification path validation policy ID, the device shall produce a CertPathValidationPolicyID fault.
	*/
	GetCertPathValidationPolicy(request *GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, error)

	GetCertPathValidationPolicyContext(ctx context.Context, request *GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, error)

	/*
	   This operation returns all certification path validation policies that are stored in the keystore on the device.
	   If no certification path validation policy is stored in the device’s keystore, an empty list is returned.
	*/
	GetAllCertPathValidationPolicies(request *GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, error)

	GetAllCertPathValidationPoliciesContext(ctx context.Context, request *GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, error)

	/*
	   This operation deletes a certification path validation policy from the keystore on the device.
	   Certification path validation policies are uniquely identified using certification path validation policy IDs. If no certification path validation policy is stored under the requested certification path validation policy ID, the device shall produce an InvalidCertPathValidationPolicyID fault.
	   If a reference exists for the requested certification path validation policy, the device shall produce a ReferenceExists fault and shall not delete the certification path validation policy.
	   After the certification path validation policy has been deleted, the device may assign its former ID to other certification path validation policies.
	*/
	DeleteCertPathValidationPolicy(request *DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, error)

	DeleteCertPathValidationPolicyContext(ctx context.Context, request *DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, error)
}

type keystore struct {
	client *soap.Client
	xaddr  string
}

func NewKeystore(client *soap.Client, xaddr string) Keystore {
	return &keystore{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *keystore) CreateRSAKeyPairContext(ctx context.Context, request *CreateRSAKeyPair) (*CreateRSAKeyPairResponse, error) {
	response := new(CreateRSAKeyPairResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/CreateRSAKeyPair", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) CreateRSAKeyPair(request *CreateRSAKeyPair) (*CreateRSAKeyPairResponse, error) {
	return service.CreateRSAKeyPairContext(
		context.Background(),
		request,
	)
}

func (service *keystore) UploadKeyPairInPKCS8Context(ctx context.Context, request *UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, error) {
	response := new(UploadKeyPairInPKCS8Response)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/UploadKeyPairInPKCS8", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) UploadKeyPairInPKCS8(request *UploadKeyPairInPKCS8) (*UploadKeyPairInPKCS8Response, error) {
	return service.UploadKeyPairInPKCS8Context(
		context.Background(),
		request,
	)
}

func (service *keystore) UploadCertificateWithPrivateKeyInPKCS12Context(ctx context.Context, request *UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, error) {
	response := new(UploadCertificateWithPrivateKeyInPKCS12Response)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/UploadCertificateWithPrivateKeyInPKCS12", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) UploadCertificateWithPrivateKeyInPKCS12(request *UploadCertificateWithPrivateKeyInPKCS12) (*UploadCertificateWithPrivateKeyInPKCS12Response, error) {
	return service.UploadCertificateWithPrivateKeyInPKCS12Context(
		context.Background(),
		request,
	)
}

func (service *keystore) GetKeyStatusContext(ctx context.Context, request *GetKeyStatus) (*GetKeyStatusResponse, error) {
	response := new(GetKeyStatusResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetKeyStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetKeyStatus(request *GetKeyStatus) (*GetKeyStatusResponse, error) {
	return service.GetKeyStatusContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetPrivateKeyStatusContext(ctx context.Context, request *GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, error) {
	response := new(GetPrivateKeyStatusResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetPrivateKeyStatus", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetPrivateKeyStatus(request *GetPrivateKeyStatus) (*GetPrivateKeyStatusResponse, error) {
	return service.GetPrivateKeyStatusContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllKeysContext(ctx context.Context, request *GetAllKeys) (*GetAllKeysResponse, error) {
	response := new(GetAllKeysResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllKeys", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllKeys(request *GetAllKeys) (*GetAllKeysResponse, error) {
	return service.GetAllKeysContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeleteKeyContext(ctx context.Context, request *DeleteKey) (*DeleteKeyResponse, error) {
	response := new(DeleteKeyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteKey", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeleteKey(request *DeleteKey) (*DeleteKeyResponse, error) {
	return service.DeleteKeyContext(
		context.Background(),
		request,
	)
}

func (service *keystore) CreatePKCS10CSRContext(ctx context.Context, request *CreatePKCS10CSR) (*CreatePKCS10CSRResponse, error) {
	response := new(CreatePKCS10CSRResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/CreatePKCS10CSR", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) CreatePKCS10CSR(request *CreatePKCS10CSR) (*CreatePKCS10CSRResponse, error) {
	return service.CreatePKCS10CSRContext(
		context.Background(),
		request,
	)
}

func (service *keystore) CreateSelfSignedCertificateContext(ctx context.Context, request *CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, error) {
	response := new(CreateSelfSignedCertificateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/CreateSelfSignedCertificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) CreateSelfSignedCertificate(request *CreateSelfSignedCertificate) (*CreateSelfSignedCertificateResponse, error) {
	return service.CreateSelfSignedCertificateContext(
		context.Background(),
		request,
	)
}

func (service *keystore) UploadCertificateContext(ctx context.Context, request *UploadCertificate) (*UploadCertificateResponse, error) {
	response := new(UploadCertificateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/UploadCertificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) UploadCertificate(request *UploadCertificate) (*UploadCertificateResponse, error) {
	return service.UploadCertificateContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetCertificateContext(ctx context.Context, request *GetCertificate) (*GetCertificateResponse, error) {
	response := new(GetCertificateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetCertificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetCertificate(request *GetCertificate) (*GetCertificateResponse, error) {
	return service.GetCertificateContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllCertificatesContext(ctx context.Context, request *GetAllCertificates) (*GetAllCertificatesResponse, error) {
	response := new(GetAllCertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllCertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllCertificates(request *GetAllCertificates) (*GetAllCertificatesResponse, error) {
	return service.GetAllCertificatesContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeleteCertificateContext(ctx context.Context, request *DeleteCertificate) (*DeleteCertificateResponse, error) {
	response := new(DeleteCertificateResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteCertificate", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeleteCertificate(request *DeleteCertificate) (*DeleteCertificateResponse, error) {
	return service.DeleteCertificateContext(
		context.Background(),
		request,
	)
}

func (service *keystore) CreateCertificationPathContext(ctx context.Context, request *CreateCertificationPath) (*CreateCertificationPathResponse, error) {
	response := new(CreateCertificationPathResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/CreateCertificationPath", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) CreateCertificationPath(request *CreateCertificationPath) (*CreateCertificationPathResponse, error) {
	return service.CreateCertificationPathContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetCertificationPathContext(ctx context.Context, request *GetCertificationPath) (*GetCertificationPathResponse, error) {
	response := new(GetCertificationPathResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetCertificationPath", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetCertificationPath(request *GetCertificationPath) (*GetCertificationPathResponse, error) {
	return service.GetCertificationPathContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllCertificationPathsContext(ctx context.Context, request *GetAllCertificationPaths) (*GetAllCertificationPathsResponse, error) {
	response := new(GetAllCertificationPathsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllCertificationPaths", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllCertificationPaths(request *GetAllCertificationPaths) (*GetAllCertificationPathsResponse, error) {
	return service.GetAllCertificationPathsContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeleteCertificationPathContext(ctx context.Context, request *DeleteCertificationPath) (*DeleteCertificationPathResponse, error) {
	response := new(DeleteCertificationPathResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteCertificationPath", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeleteCertificationPath(request *DeleteCertificationPath) (*DeleteCertificationPathResponse, error) {
	return service.DeleteCertificationPathContext(
		context.Background(),
		request,
	)
}

func (service *keystore) UploadPassphraseContext(ctx context.Context, request *UploadPassphrase) (*UploadPassphraseResponse, error) {
	response := new(UploadPassphraseResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/UploadPassphrase", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) UploadPassphrase(request *UploadPassphrase) (*UploadPassphraseResponse, error) {
	return service.UploadPassphraseContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllPassphrasesContext(ctx context.Context, request *GetAllPassphrases) (*GetAllPassphrasesResponse, error) {
	response := new(GetAllPassphrasesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllPassphrases", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllPassphrases(request *GetAllPassphrases) (*GetAllPassphrasesResponse, error) {
	return service.GetAllPassphrasesContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeletePassphraseContext(ctx context.Context, request *DeletePassphrase) (*DeletePassphraseResponse, error) {
	response := new(DeletePassphraseResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeletePassphrase", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeletePassphrase(request *DeletePassphrase) (*DeletePassphraseResponse, error) {
	return service.DeletePassphraseContext(
		context.Background(),
		request,
	)
}

func (service *keystore) UploadCRLContext(ctx context.Context, request *UploadCRL) (*UploadCRLResponse, error) {
	response := new(UploadCRLResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/UploadCRL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) UploadCRL(request *UploadCRL) (*UploadCRLResponse, error) {
	return service.UploadCRLContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetCRLContext(ctx context.Context, request *GetCRL) (*GetCRLResponse, error) {
	response := new(GetCRLResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetCRL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetCRL(request *GetCRL) (*GetCRLResponse, error) {
	return service.GetCRLContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllCRLsContext(ctx context.Context, request *GetAllCRLs) (*GetAllCRLsResponse, error) {
	response := new(GetAllCRLsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllCRLs", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllCRLs(request *GetAllCRLs) (*GetAllCRLsResponse, error) {
	return service.GetAllCRLsContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeleteCRLContext(ctx context.Context, request *DeleteCRL) (*DeleteCRLResponse, error) {
	response := new(DeleteCRLResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteCRL", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeleteCRL(request *DeleteCRL) (*DeleteCRLResponse, error) {
	return service.DeleteCRLContext(
		context.Background(),
		request,
	)
}

func (service *keystore) CreateCertPathValidationPolicyContext(ctx context.Context, request *CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, error) {
	response := new(CreateCertPathValidationPolicyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/CreateCertPathValidationPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) CreateCertPathValidationPolicy(request *CreateCertPathValidationPolicy) (*CreateCertPathValidationPolicyResponse, error) {
	return service.CreateCertPathValidationPolicyContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetCertPathValidationPolicyContext(ctx context.Context, request *GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, error) {
	response := new(GetCertPathValidationPolicyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetCertPathValidationPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetCertPathValidationPolicy(request *GetCertPathValidationPolicy) (*GetCertPathValidationPolicyResponse, error) {
	return service.GetCertPathValidationPolicyContext(
		context.Background(),
		request,
	)
}

func (service *keystore) GetAllCertPathValidationPoliciesContext(ctx context.Context, request *GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, error) {
	response := new(GetAllCertPathValidationPoliciesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllCertPathValidationPolicies", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) GetAllCertPathValidationPolicies(request *GetAllCertPathValidationPolicies) (*GetAllCertPathValidationPoliciesResponse, error) {
	return service.GetAllCertPathValidationPoliciesContext(
		context.Background(),
		request,
	)
}

func (service *keystore) DeleteCertPathValidationPolicyContext(ctx context.Context, request *DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, error) {
	response := new(DeleteCertPathValidationPolicyResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteCertPathValidationPolicy", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *keystore) DeleteCertPathValidationPolicy(request *DeleteCertPathValidationPolicy) (*DeleteCertPathValidationPolicyResponse, error) {
	return service.DeleteCertPathValidationPolicyContext(
		context.Background(),
		request,
	)
}

type TLSServer interface {

	/*
		This operation assigns a key pair and certificate along with a certification path (certificate chain) to the TLS server on the device.
		The TLS server shall use this information for key exchange during the TLS handshake, particularly for constructing server certificate messages as specified in RFC 4346 and RFC 2246.

		Certification paths are identified by their certification path IDs in the keystore. The first certificate in the certification path must be the TLS server certificate.
		Since each certificate has exactly one associated key pair, a reference to the key pair that is associated with the server certificate is not supplied explicitly.
		Devices shall obtain the private key or results of operations under the private key by suitable internal interaction with the keystore.
		If a device chooses to perform a TLS key exchange based on the supplied certification path,  it shall use the key pair that is associated with the server certificate for
		key exchange and transmit the certification path to TLS clients as-is, i.e., the device shall not check conformance of the certification path to RFC 4346 norRFC 2246.
		In order to use the server certificate during the TLS handshake, the corresponding private key is required.
		Therefore, if the key pair that is associated with the server certificate, i.e., the first certificate in the certification path, does not have an associated private key,
		the NoPrivateKey fault is produced and the certification path is not associated to the TLS server.
		A TLS server may present different certification paths to different clients during the TLS handshake instead of presenting the same certification path to all clients.
		Therefore more than one certification path may be assigned to the TLS server.
		If the maximum number of certification paths that may be assigned to the TLS server simultaneously is reached, the device shall generate a MaximumNumberOfCertificationPathsReached
		fault and the requested certification path shall not be assigned to the TLS server.
	*/
	AddServerCertificateAssignment(request *AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, error)

	AddServerCertificateAssignmentContext(ctx context.Context, request *AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, error)

	/*
		This operation removes a key pair and certificate assignment (including certification path) to the TLS server on the device.
		Certification paths are identified using certification path IDs. If the supplied certification path ID is not associated to the TLS server, an InvalidArgVal fault is produced.
	*/
	RemoveServerCertificateAssignment(request *RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, error)

	RemoveServerCertificateAssignmentContext(ctx context.Context, request *RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, error)

	/*
		This operation replaces an existing key pair and certificate assignment to the TLS server on the device by a new key pair and certificate assignment (including certification paths).

		After the replacement, the TLS server shall use the new certificate and certification path exactly in those cases in which it would have used the old certificate and certification path.
		Therefore, especially in the case that several server certificates are assigned to the TLS server, clients that wish to replace an old certificate assignment by a new assignment
		should use this operation instead of a combination of the Add TLS Server Certificate Assignment and the Remove TLS Server Certificate Assignment operations.

		Certification paths are identified using certification path IDs. If the supplied old certification path ID is not associated to the TLS server, or no certification path exists
		under the new certification path ID, the corresponding InvalidArgVal faults are produced and the associations are unchanged.
		The first certificate in the new certification path must be the TLS server certificate.
		Since each certificate has exactly one associated key pair, a reference to the key pair that is associated with the new server certificate is not supplied explicitly.
		Devices shall obtain the private key or results of operations under the private key by suitable internal interaction with the keystore.
		If a device chooses to perform a TLS key exchange based on the new certification path,  it shall use the key pair that is associated with the server certificate
		for key exchange and transmit the certification path to TLS clients as-is, i.e., the device shall not check conformance of the certification path to RFC 4346 norRFC 2246.
		In order to use the server certificate during the TLS handshake, the corresponding private key is required.
		Therefore, if the key pair that is associated with the server certificate, i.e., the first certificate in the certification path, does not have an associated private key,
		the NoPrivateKey fault is produced and the certification path is not associated to the TLS server.
	*/
	ReplaceServerCertificateAssignment(request *ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, error)

	ReplaceServerCertificateAssignmentContext(ctx context.Context, request *ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, error)

	/*
		This operation sets the version(s) of TLS which the device shall use.  Valid values are taken from the TLSServerSupported capability.
		A client initiates a TLS session by sending a ClientHello with the hightest TLS version it supports.  This suggests to the server that the client can accept any TLS version up to and including that version.
		The server then chooses the TLS version to use.  This is generally the highest TLS version the server supports that is within the range of the client.  For example, if a ClientHello indicates TLS version 1.1, the server can proceed with TLS 1.0 or TLS 1.1.
		In the event that an ONVIF installation wishes to disable certain version(s) of TLS, it may do so with this operation.  For example, to disable TLS 1.0 on a device signaling support for TLS versions 1.0, 1.1, and 1.2, the enabled version list may be set to "1.1 1.2", omitting 1.0.  If a client then attempts to connect with a ClientHello containing TLS 1.0, the server shall send a "protocol_version" alert message and close the connection.  This handshake indicates to the client that TLS 1.0 is not supported by the server.  The client must try again with a higher TLS version suggestion.
		An empty list is not permitted.  Disabling all versions of TLS is not the intent of this operation.  See AddServerCertificateAssignment and RemoveServerCertificateAssignment.
	*/
	SetEnabledTLSVersions(request *SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, error)

	SetEnabledTLSVersionsContext(ctx context.Context, request *SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, error)

	/*
		This operation retrieves the version(s) of TLS which are currently enabled on the device.
	*/
	GetEnabledTLSVersions(request *GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, error)

	GetEnabledTLSVersionsContext(ctx context.Context, request *GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, error)

	/*
		This operation returns the IDs of all key pairs and certificates (including certification paths) that are assigned to the TLS server on the device.
		This operation may be used, e.g., if a client lost track of the certification path assignments on the device.
		If no certification path is assigned to the TLS server, an empty list is returned.
	*/
	GetAssignedServerCertificates(request *GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, error)

	GetAssignedServerCertificatesContext(ctx context.Context, request *GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, error)

	/*
	   This operation activates or deactivates TLS client authentication for the TLS server on the device.
	   The TLS server on the device shall require client authentication if and only if clientAuthenticationRequired is set to true.
	   If TLS client authentication is requested to be enabled and no certification path validation policy is assigned to the TLS server, the device shall return an EnablingTLSClientAuthenticationFailed fault and shall not enable TLS client authentication.
	   The device shall execute this command regardless of the TLS enabled/disabled state configured in the ONVIF Device Management Service.
	*/
	SetClientAuthenticationRequired(request *SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, error)

	SetClientAuthenticationRequiredContext(ctx context.Context, request *SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, error)

	/*
	   This operation returns whether TLS client authentication is active.
	*/
	GetClientAuthenticationRequired(request *GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, error)

	GetClientAuthenticationRequiredContext(ctx context.Context, request *GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, error)

	/*
	   This operation assigns a certification path validation policy to the TLS server on the device. The TLS server shall enforce the policy when authenticating TLS clients and consider a client authentic if and only if the algorithm returns valid.
	   If no certification path validation policy is stored under the requested CertPathValidationPolicyID, the device shall produce a CertPathValidationPolicyID fault.
	   A TLS server may use different certification path validation policies to authenticate clients. Therefore more than one certification path validation policy may be assigned to the TLS server. If the maximum number of certification path validation policies that may be assigned to the  TLS server simultaneously is reached, the device shall produce a MaximumNumberOfTLSCertPathValidationPoliciesReached fault and shall not assign the requested certification path validation policy to the TLS server.
	*/
	AddCertPathValidationPolicyAssignment(request *AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, error)

	AddCertPathValidationPolicyAssignmentContext(ctx context.Context, request *AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, error)

	/*
	   This operation removes a certification path validation policy assignment from the TLS server on the device.
	   If the certification path validation policy identified by the requested CertPathValidationPolicyID is not associated to the TLS server, the device shall produce a CertPathValidationPolicy fault.
	*/
	RemoveCertPathValidationPolicyAssignment(request *RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, error)

	RemoveCertPathValidationPolicyAssignmentContext(ctx context.Context, request *RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, error)

	/*
	   This operation replaces a certification path validation policy assignment to the TLS server on the device with another certification path validation policy assignment.
	   If the certification path validation policy identified by the requested OldCertPathValidationPolicyID is not associated to the TLS server, the device shall produce an OldCertPathValidationPolicyID fault and shall not associate the certification path validation policy identified by the NewCertPathValidationPolicyID to the TLS server.
	   If no certification path validation policy exists under the requested NewCertPathValidationPolicyID in the device’s keystore, the device shall produce a NewCertPathValidationPolicyID fault and shall not remove the association of the old certification path validation policy to the TLS server.
	*/
	ReplaceCertPathValidationPolicyAssignment(request *ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, error)

	ReplaceCertPathValidationPolicyAssignmentContext(ctx context.Context, request *ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, error)

	/*
	   This operation returns the IDs of all certification path validation policies that are assigned to the TLS server on the device.
	*/
	GetAssignedCertPathValidationPolicies(request *GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, error)

	GetAssignedCertPathValidationPoliciesContext(ctx context.Context, request *GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, error)
}

type tLSServer struct {
	client *soap.Client
	xaddr  string
}

func NewTLSServer(client *soap.Client, xaddr string) TLSServer {
	return &tLSServer{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *tLSServer) AddServerCertificateAssignmentContext(ctx context.Context, request *AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, error) {
	response := new(AddServerCertificateAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/AddServerCertificateAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) AddServerCertificateAssignment(request *AddServerCertificateAssignment) (*AddServerCertificateAssignmentResponse, error) {
	return service.AddServerCertificateAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) RemoveServerCertificateAssignmentContext(ctx context.Context, request *RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, error) {
	response := new(RemoveServerCertificateAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/RemoveServerCertificateAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) RemoveServerCertificateAssignment(request *RemoveServerCertificateAssignment) (*RemoveServerCertificateAssignmentResponse, error) {
	return service.RemoveServerCertificateAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) ReplaceServerCertificateAssignmentContext(ctx context.Context, request *ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, error) {
	response := new(ReplaceServerCertificateAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/ReplaceServerCertificateAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) ReplaceServerCertificateAssignment(request *ReplaceServerCertificateAssignment) (*ReplaceServerCertificateAssignmentResponse, error) {
	return service.ReplaceServerCertificateAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) SetEnabledTLSVersionsContext(ctx context.Context, request *SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, error) {
	response := new(SetEnabledTLSVersionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/SetEnabledTLSVersions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) SetEnabledTLSVersions(request *SetEnabledTLSVersions) (*SetEnabledTLSVersionsResponse, error) {
	return service.SetEnabledTLSVersionsContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) GetEnabledTLSVersionsContext(ctx context.Context, request *GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, error) {
	response := new(GetEnabledTLSVersionsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetEnabledTLSVersions", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) GetEnabledTLSVersions(request *GetEnabledTLSVersions) (*GetEnabledTLSVersionsResponse, error) {
	return service.GetEnabledTLSVersionsContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) GetAssignedServerCertificatesContext(ctx context.Context, request *GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, error) {
	response := new(GetAssignedServerCertificatesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAssignedServerCertificates", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) GetAssignedServerCertificates(request *GetAssignedServerCertificates) (*GetAssignedServerCertificatesResponse, error) {
	return service.GetAssignedServerCertificatesContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) SetClientAuthenticationRequiredContext(ctx context.Context, request *SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, error) {
	response := new(SetClientAuthenticationRequiredResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/SetClientAuthenticationRequired", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) SetClientAuthenticationRequired(request *SetClientAuthenticationRequired) (*SetClientAuthenticationRequiredResponse, error) {
	return service.SetClientAuthenticationRequiredContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) GetClientAuthenticationRequiredContext(ctx context.Context, request *GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, error) {
	response := new(GetClientAuthenticationRequiredResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetClientAuthenticationRequired", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) GetClientAuthenticationRequired(request *GetClientAuthenticationRequired) (*GetClientAuthenticationRequiredResponse, error) {
	return service.GetClientAuthenticationRequiredContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) AddCertPathValidationPolicyAssignmentContext(ctx context.Context, request *AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, error) {
	response := new(AddCertPathValidationPolicyAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/AddCertPathValidationPolicyAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) AddCertPathValidationPolicyAssignment(request *AddCertPathValidationPolicyAssignment) (*AddCertPathValidationPolicyAssignmentResponse, error) {
	return service.AddCertPathValidationPolicyAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) RemoveCertPathValidationPolicyAssignmentContext(ctx context.Context, request *RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, error) {
	response := new(RemoveCertPathValidationPolicyAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/RemoveCertPathValidationPolicyAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) RemoveCertPathValidationPolicyAssignment(request *RemoveCertPathValidationPolicyAssignment) (*RemoveCertPathValidationPolicyAssignmentResponse, error) {
	return service.RemoveCertPathValidationPolicyAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) ReplaceCertPathValidationPolicyAssignmentContext(ctx context.Context, request *ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, error) {
	response := new(ReplaceCertPathValidationPolicyAssignmentResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/ReplaceCertPathValidationPolicyAssignment", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) ReplaceCertPathValidationPolicyAssignment(request *ReplaceCertPathValidationPolicyAssignment) (*ReplaceCertPathValidationPolicyAssignmentResponse, error) {
	return service.ReplaceCertPathValidationPolicyAssignmentContext(
		context.Background(),
		request,
	)
}

func (service *tLSServer) GetAssignedCertPathValidationPoliciesContext(ctx context.Context, request *GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, error) {
	response := new(GetAssignedCertPathValidationPoliciesResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAssignedCertPathValidationPolicies", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *tLSServer) GetAssignedCertPathValidationPolicies(request *GetAssignedCertPathValidationPolicies) (*GetAssignedCertPathValidationPoliciesResponse, error) {
	return service.GetAssignedCertPathValidationPoliciesContext(
		context.Background(),
		request,
	)
}

type Dot1X interface {

	/*
		(to be written)
	*/
	AddDot1XConfiguration(request *AddDot1XConfiguration) (*AddDot1XConfigurationResponse, error)

	AddDot1XConfigurationContext(ctx context.Context, request *AddDot1XConfiguration) (*AddDot1XConfigurationResponse, error)

	/*
		(to be written)
	*/
	GetAllDot1XConfigurations(request *GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, error)

	GetAllDot1XConfigurationsContext(ctx context.Context, request *GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, error)

	/*
		(to be written)
	*/
	GetDot1XConfiguration(request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error)

	GetDot1XConfigurationContext(ctx context.Context, request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error)

	/*
		(to be written)
	*/
	DeleteDot1XConfiguration(request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error)

	DeleteDot1XConfigurationContext(ctx context.Context, request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error)

	/*
		(to be written)
	*/
	SetNetworkInterfaceDot1XConfiguration(request *SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, error)

	SetNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, error)

	/*
		(to be written)
	*/
	GetNetworkInterfaceDot1XConfiguration(request *GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, error)

	GetNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, error)

	/*
		(to be written)
	*/
	DeleteNetworkInterfaceDot1XConfiguration(request *DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, error)

	DeleteNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, error)
}

type dot1X struct {
	client *soap.Client
	xaddr  string
}

func NewDot1X(client *soap.Client, xaddr string) Dot1X {
	return &dot1X{
		client: client,
		xaddr:  xaddr,
	}
}

func (service *dot1X) AddDot1XConfigurationContext(ctx context.Context, request *AddDot1XConfiguration) (*AddDot1XConfigurationResponse, error) {
	response := new(AddDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/AddDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) AddDot1XConfiguration(request *AddDot1XConfiguration) (*AddDot1XConfigurationResponse, error) {
	return service.AddDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) GetAllDot1XConfigurationsContext(ctx context.Context, request *GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, error) {
	response := new(GetAllDot1XConfigurationsResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetAllDot1XConfigurations", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) GetAllDot1XConfigurations(request *GetAllDot1XConfigurations) (*GetAllDot1XConfigurationsResponse, error) {
	return service.GetAllDot1XConfigurationsContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) GetDot1XConfigurationContext(ctx context.Context, request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error) {
	response := new(GetDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) GetDot1XConfiguration(request *GetDot1XConfiguration) (*GetDot1XConfigurationResponse, error) {
	return service.GetDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) DeleteDot1XConfigurationContext(ctx context.Context, request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error) {
	response := new(DeleteDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) DeleteDot1XConfiguration(request *DeleteDot1XConfiguration) (*DeleteDot1XConfigurationResponse, error) {
	return service.DeleteDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) SetNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, error) {
	response := new(SetNetworkInterfaceDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/SetNetworkInterfaceDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) SetNetworkInterfaceDot1XConfiguration(request *SetNetworkInterfaceDot1XConfiguration) (*SetNetworkInterfaceDot1XConfigurationResponse, error) {
	return service.SetNetworkInterfaceDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) GetNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, error) {
	response := new(GetNetworkInterfaceDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/GetNetworkInterfaceDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) GetNetworkInterfaceDot1XConfiguration(request *GetNetworkInterfaceDot1XConfiguration) (*GetNetworkInterfaceDot1XConfigurationResponse, error) {
	return service.GetNetworkInterfaceDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

func (service *dot1X) DeleteNetworkInterfaceDot1XConfigurationContext(ctx context.Context, request *DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, error) {
	response := new(DeleteNetworkInterfaceDot1XConfigurationResponse)
	err := service.client.CallContext(ctx, service.xaddr, "http://www.onvif.org/ver10/advancedsecurity/wsdl/DeleteNetworkInterfaceDot1XConfiguration", request, response)
	if err != nil {
		return nil, err
	}

	return response, nil
}

func (service *dot1X) DeleteNetworkInterfaceDot1XConfiguration(request *DeleteNetworkInterfaceDot1XConfiguration) (*DeleteNetworkInterfaceDot1XConfigurationResponse, error) {
	return service.DeleteNetworkInterfaceDot1XConfigurationContext(
		context.Background(),
		request,
	)
}

type AnyURI string
type Duration string
type QName string
type NCName string
type NonNegativeInteger int64
type PositiveInteger int64
type NonPositiveInteger int64
type AnySimpleType string
type String string
