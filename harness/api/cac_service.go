package api

// func (c *ConfigAsCodeClient) GetServiceById(applicationId string, serviceId string) (*cac.Service, error) {
// 	item, err := c.GetDirectoryItemContent("services", serviceId, applicationId)
// 	if err != nil {
// 		return nil, err
// 	}

// 	// Item not found
// 	if item == nil {
// 		return nil, nil
// 	}

// 	svc := &cac.Service{}
// 	err = item.ParseYamlContent(svc)
// 	if err != nil {
// 		return nil, err
// 	}

// 	svc.ApplicationId = applicationId

// 	return svc, nil
// }

// func (c *ConfigAsCodeClient) DeleteService(applicationId string, serviceId string) error {

// 	app, err := c.ApiClient.Applications().GetApplicationById(applicationId)
// 	if err != nil {
// 		return err
// 	}

// 	if app == nil {
// 		return fmt.Errorf("could not find application by id: '%s'", applicationId)
// 	}

// 	svc, err := c.GetServiceById(applicationId, serviceId)
// 	if err != nil {
// 		return err
// 	}

// 	filePath := fmt.Sprintf("Setup/Applications/%s/Services/%s/Index.yaml", app.Name, svc.Name)

// 	return c.DeleteEntities([]string{filePath})
// }
