package services

type tmtReportServicesDAO interface {
	Get(date string) (string, error)
}

type tmtReportServices struct {
	dao tmtReportServicesDAO
}

func NewTMTReportServices(dao tmtReportServicesDAO) *tmtReportServices {

	return &tmtReportServices{dao: dao}
}

func (s tmtReportServices) Get(date string) (string, error) {

	return s.dao.Get(date)
}
