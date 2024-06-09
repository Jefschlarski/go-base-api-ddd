package repositories

import (
	"api/src/application/interfaces"
)

type AddressRepository interface {
	interfaces.CreateAddress
	interfaces.GetAddressByUserId
	interfaces.UpdateAddress
	interfaces.GetAddress
	interfaces.GetAllAddresses
	interfaces.DeleteAddress
}
