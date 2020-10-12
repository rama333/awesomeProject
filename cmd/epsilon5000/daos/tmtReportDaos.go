package daos

type tmtReportDaos struct{}

func (t tmtReportDaos) New(date string) *tmtReportDaos {

	return &tmtReportDaos{}
}

func Get(date string) (string, error) {

	return "", nil
}
