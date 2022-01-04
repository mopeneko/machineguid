package machineguid

import (
	"golang.org/x/sys/windows/registry"
	"golang.org/x/xerrors"
)

// GetMachineGUID returns Machine GUID, a UUID which should in theory be unique to the machine
func GetMachineGUID() (string, error) {
	key, err := registry.OpenKey(registry.LOCAL_MACHINE, `SOFTWARE\Microsoft\Cryptography`, registry.QUERY_VALUE)
	if err != nil {
		return "", xerrors.Errorf("failed to open key: %w", err)
	}

	defer key.Close()

	machineGUID, _, err := key.GetStringValue("MachineGUID")
	if err != nil {
		return "", xerrors.Errorf("failed to get MachineGUID value: %w", err)
	}

	return machineGUID, nil
}
