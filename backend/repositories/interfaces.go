package repositories

import (
	"time"

	"github.com/escoteirando/mappa-proxy/backend/domain"
	"github.com/escoteirando/mappa-proxy/backend/domain/entities"
	"github.com/escoteirando/mappa-proxy/backend/domain/responses"
	"gorm.io/gorm"
)

type IRepository interface {
	IsValidConnectionString(connectionString string) bool
	GetDBFunc() *gorm.DB
	CreateRepository(connectionString string) (IRepository, error)
	GetName() string

	GetLogins() (logins []*domain.LoginData, err error) //

	SetLogin(username string, password string, loginResponse responses.MappaLoginResponse, last_login time.Time) error //
	// GetLogin(username string, password string) (loginResponse *responses.MappaLoginResponse, err error)
	DeleteLogin(username string) error //
	// SaveData() error
	// loadData() error

	// GetLastLogin() (username string, timestamp time.Time)
	// SetLastLogin(username string, timestamp time.Time) error
	// GetUserCount() int

	SetEscotista(escotista *entities.Escotista) error
	GetEscotista(userId int) (escotista *entities.Escotista, err error)

	SetAssociado(associado *entities.Associado) error
	GetAssociado(codigoAssociado int) (associado *entities.Associado, err error)

	SetGrupo(grupo *entities.Grupo) error
	GetGrupo(codigoGrupo int) (grupo *entities.Grupo, err error)

	SetSecao(secao *entities.Secao) error
	GetSecao(codigoSecao int, codigoRegiao string) (secao *entities.Secao, err error)

	SetSubSecao(subsecao *entities.SubSecao) error
	GetSubSecao(codigoSubSecao int) (subsecao *entities.SubSecao, err error)
	GetSubSecoes(codigoSecao int) (subsecoes []*entities.SubSecao, err error)
	GetSubSecaoAssociados(codigoSubSecao int) (associados []entities.Associado, err error)
	UnlinkSubSecaoAssociados(subSecao *entities.SubSecao) error

	SetAssociadoSecao(codigoAssociado int, codigosSubSecao ...int) error
	GetAssociadoSecoes(codigoAssociado int) (secoes []*entities.Secao, err error)
	SetAssociadoSubSecao(codigoAssociado int, codigoSubSecao int) error
	GetAssociadoSubSecao(codigoAssociado int) (subsecao *entities.SubSecao, err error)

	SetKeyValue(key, value string, timeToLive time.Duration) error
	GetKeyValue(key, defaultValue string) string
	DeleteKey(key string) error
	DeleteKeys(keyPrefix string) error

	UpdateMappaProgressoes(progressoesResponse []*responses.MappaProgressaoResponse) error
	GetProgressoes(domain.MappaRamoEnum) ([]*responses.MappaProgressaoResponse, error)

	UpdateMappaMarcacoes(marcacoes []*entities.MappaMarcacao) error
	GetMarcacoes(codigoSecao int) ([]*entities.MappaMarcacao, error)

	UpdateMappaEspecialidades(especialidades []*entities.MappaEspecialidade) error
	GetEspecialidades() ([]*entities.MappaEspecialidade, error)
	GetEspecialidade(codEspecialidade int) (*entities.MappaEspecialidade, error)

	UpdateMappaConquistas(conquistas []*entities.MappaConquista) error
	GetConquistas(codigoSecao int, since time.Time) ([]*entities.MappaConquista, error)

	GetCounts() (map[string]int, error)
}
