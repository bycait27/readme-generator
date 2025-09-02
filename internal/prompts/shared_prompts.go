package prompts

import (
	"github.com/bycait27/readme-generator/internal/models"
	"github.com/bycait27/readme-generator/internal/validation"
)

// getting started info
func PromptGettingStartedInfo() (*models.GettingStarted, error) {
	// prompt for prerequisites
	prereqs, err := promptStringList("prerequisites", false)
	if err != nil {
		return nil, err
	}

	// prompt for environment setup
	envSetup, err := promptStringList("environment setup steps", false)
	if err != nil {
		return nil, err
	}

	// prompt for run commands
	runCommands, err := promptStringList("Run Commands", true)
	if err != nil {
		return nil, err
	}

	// prompt for notes
	notes, err := promptOptionalStringPointer("any additional notes", 500)
	if err != nil {
		return nil, err
	}

	// create GettingStartedInfo
	gettingStartedInfo := &models.GettingStarted{
		Prerequisites: prereqs,  // optional
		EnvSetup:      envSetup, // optional
		RunCommands:   runCommands,
		Notes:         notes, // optional
	}

	// final validation
	if err := validation.ValidateStruct(gettingStartedInfo); err != nil {
		return nil, err
	}

	return gettingStartedInfo, nil
}

// environmental variables info
func PromptEnvVarInfo() (*models.EnvVar, error) {
	// prompt for name
	name, err := promptRequiredText("Environment variable name", 1, 50)
	if err != nil {
		return nil, err
	}

	// prompt for description
	description, err := promptRequiredText("Environment variable description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt for required
	required, err := promptRequiredBoolean("Is this environment variable required?")
	if err != nil {
		return nil, err
	}

	// prompt for default
	defaultValue, err := promptOptionalText("a default value", 100)
	if err != nil {
		return nil, err
	}

	// prompt for example
	example, err := promptRequiredText("Provide an example", 1, 100)
	if err != nil {
		return nil, err
	}

	// create envVarInfo
	envVarInfo := &models.EnvVar{
		Name:        name,
		Description: description,
		Required:    required,
		Default:     defaultValue, // optional
		Example:     example,
	}

	// final validation
	if err := validation.ValidateStruct(envVarInfo); err != nil {
		return nil, err
	}

	return envVarInfo, nil
}

// api documentation info
func PromptAPIDocsInfo() (*models.APIDocs, error) {
	// prompt base url
	baseURL, err := promptURL("Base URL", true)
	if err != nil {
		return nil, err
	}

	// prompt authentication
	auth, err := promptRequiredText("Authentication type", 1, 50)
	if err != nil {
		return nil, err
	}

	// prompt endpoints
	endpoint, err := PromptEndpointInfo()
	if err != nil {
		return nil, err
	}
	endpoints := []models.Endpoint{}
	if endpoint != nil {
		endpoints = append(endpoints, *endpoint)
	}

	// prompt error handling
	var errHandling *models.ErrorHandling
	wantErrHandling, err := promptYesNo("Do you want to include error handling information?")
	if err != nil {
		return nil, err
	}
	if wantErrHandling {
		errHandling, err = PromptErrorHandlingInfo()
		if err != nil {
			return nil, err
		}
	}

	// create APIDocsInfo
	apiDocsInfo := &models.APIDocs{
		BaseURL:        baseURL,
		Authentication: auth,
		Endpoints:      endpoints,
		ErrorHandling:  errHandling, // optional
	}

	// final validation
	if err := validation.ValidateStruct(apiDocsInfo); err != nil {
		return nil, err
	}

	return apiDocsInfo, nil
}

// endpoint info
func PromptEndpointInfo() (*models.Endpoint, error) {
	// prompt method
	method, err := promptFromOptions("Choose endpoint method", []string{"GET", "POST", "PUT", "DELETE", "PATCH", "OPTIONS", "HEAD"})
	if err != nil {
		return nil, err
	}

	// prompt path
	path, err := promptRequiredText("Endpoint path", 1, 50)
	if err != nil {
		return nil, err
	}

	// prompt description
	description, err := promptRequiredText("Endpoint description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt parameters
	var parameters []models.Parameter
	wantParameters, err := promptYesNo("Do you want to include any parameters?")
	if err != nil {
		return nil, err
	}
	if wantParameters {
		param, err := PromptParameterInfo()
		if err != nil {
			return nil, err
		}
		if param != nil {
			parameters = append(parameters, *param)
		}
	}

	// prompt response
	response, err := promptRequiredText("Endpoint response", 1, 200)
	if err != nil {
		return nil, err
	}

	// create EndpointInfo
	endpointInfo := &models.Endpoint{
		Method:      method,
		Path:        path,
		Description: description,
		Parameters:  parameters, // optional
		Response:    response,
	}

	// final validation
	if err := validation.ValidateStruct(endpointInfo); err != nil {
		return nil, err
	}

	return endpointInfo, nil
}

// parameter info
func PromptParameterInfo() (*models.Parameter, error) {
	// prompt name
	name, err := promptRequiredText("Parameter name", 1, 50)
	if err != nil {
		return nil, err
	}

	// prompt type
	typeValue, err := promptFromOptions("Choose a parameter type", []string{"string", "int", "bool"})
	if err != nil {
		return nil, err
	}

	// prompt required
	required, err := promptRequiredBoolean("Is this parameter required?")
	if err != nil {
		return nil, err
	}

	// prompt description
	description, err := promptRequiredText("Parameter description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt example
	example, err := promptRequiredText("Parameter example", 1, 100)
	if err != nil {
		return nil, err
	}

	// create ParameterInfo
	parameterInfo := &models.Parameter{
		Name:        name,
		Type:        typeValue,
		Required:    required,
		Description: description,
		Example:     example,
	}

	// final validation
	if err := validation.ValidateStruct(parameterInfo); err != nil {
		return nil, err
	}

	return parameterInfo, nil
}

// error handling info
func PromptErrorHandlingInfo() (*models.ErrorHandling, error) {
	// prompt format
	format, err := promptFromOptions("Choose the type of error handling", []string{"json", "xml", "plain"})
	if err != nil {
		return nil, err
	}

	// prompt status codes
	var statusCodes []models.StatusCode
	wantStatusCodes, err := promptYesNo("Do you want to include the status codes used?")
	if err != nil {
		return nil, err
	}
	if wantStatusCodes {
		statusCode, err := PromptStatusCodeInfo()
		if err != nil {
			return nil, err
		}
		if statusCode != nil {
			statusCodes = append(statusCodes, *statusCode)
		}
	}

	// prompt error response
	var errorResponses *models.ErrorResponse
	wantErrorResponses, err := promptYesNo("Do you want to include error responses?")
	if err != nil {
		return nil, err
	}
	if wantErrorResponses {
		errorResponses, err = PromptErrorResponseInfo()
		if err != nil {
			return nil, err
		}
	}

	// prompt common errors
	var commonErrors []models.CommonError
	wantCommonErrors, err := promptYesNo("Do you want to include any common errors?")
	if err != nil {
		return nil, err
	}
	if wantCommonErrors {
		commonError, err := PromptCommonErrorInfo()
		if err != nil {
			return nil, err
		}
		if commonError != nil {
			commonErrors = append(commonErrors, *commonError)
		}
	}

	// prompt validation rules
	valRules, err := promptOptionalText("error handling validation rules", 500)
	if err != nil {
		return nil, err
	}

	// create ErrorHandlingInfo
	errorHandlingInfo := &models.ErrorHandling{
		Format:          format,
		StatusCodes:     statusCodes,     // optional
		ErrorResponse:   *errorResponses, // optional
		CommonErrors:    commonErrors,    // optional
		ValidationRules: valRules,        // optional
	}

	// final validation
	if err := validation.ValidateStruct(errorHandlingInfo); err != nil {
		return nil, err
	}

	return errorHandlingInfo, nil
}

// status code info
func PromptStatusCodeInfo() (*models.StatusCode, error) {
	// prompt code
	code, err := promptRequiredInt("Status code", 100, 600)
	if err != nil {
		return nil, err
	}

	// prompt description
	description, err := promptRequiredText("Status code description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt example
	example, err := promptOptionalText("status code example", 200)
	if err != nil {
		return nil, err
	}

	// create StatusCodeInfo
	statusCodeInfo := &models.StatusCode{
		Code:        code,
		Description: description,
		Example:     example, // optional
	}

	// final validation
	if err := validation.ValidateStruct(statusCodeInfo); err != nil {
		return nil, err
	}

	return statusCodeInfo, nil
}

// error response info
func PromptErrorResponseInfo() (*models.ErrorResponse, error) {
	// prompt structure
	structure, err := promptRequiredText("Error response structure", 1, 50)
	if err != nil {
		return nil, err
	}

	// prompt example
	example, err := promptRequiredText("Example error response", 5, 200)
	if err != nil {
		return nil, err
	}

	// create ErrorResponseInfo
	errorResponseInfo := &models.ErrorResponse{
		Structure: structure,
		Example:   example,
	}

	// final validation
	if err := validation.ValidateStruct(errorResponseInfo); err != nil {
		return nil, err
	}

	return errorResponseInfo, nil
}

// common error info
func PromptCommonErrorInfo() (*models.CommonError, error) {
	// prompt code
	code, err := promptRequiredInt("Common error code", 100, 600)
	if err != nil {
		return nil, err
	}

	// prompt message
	message, err := promptRequiredText("Common error message", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt description
	description, err := promptRequiredText("Common error description", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt solution
	solution, err := promptRequiredText("Common error solution", 5, 500)
	if err != nil {
		return nil, err
	}

	// create CommonErrorInfo
	commonErrorInfo := &models.CommonError{
		Code:        code,
		Message:     message,
		Description: description,
		Solution:    solution,
	}

	// final validation
	if err := validation.ValidateStruct(commonErrorInfo); err != nil {
		return nil, err
	}

	return commonErrorInfo, nil
}

// database info
func PromptDatabaseInfo() (*models.Database, error) {
	// prompt type
	dbType, err := promptFromOptions("Choose a database type", []string{"PostgreSQL", "MySQL", "MongoDB", "Redis", "SQLite", "Cassandra", "MariaDB", "DynamoDB", "Firebase"})
	if err != nil {
		return nil, err
	}

	// prompt schema
	schema, err := promptRequiredText("Database schema", 5, 200)
	if err != nil {
		return nil, err
	}

	// prompt migrations
	migrations, err := promptOptionalText("database migrations", 500)
	if err != nil {
		return nil, err
	}

	// prompt seed data
	seedData, err := promptOptionalText("database seed data", 500)
	if err != nil {
		return nil, err
	}

	// create DatabaseInfo
	databaseInfo := &models.Database{
		Type:       dbType,
		Schema:     schema,
		Migrations: migrations, // optional
		SeedData:   seedData,   // optional
	}

	// final validation
	if err := validation.ValidateStruct(databaseInfo); err != nil {
		return nil, err
	}

	return databaseInfo, nil
}

// testing info
func PromptTestingInfo() (*models.Testing, error) {
	// prompt test command
	testCmd, err := promptRequiredText("Run tests command", 1, 200)
	if err != nil {
		return nil, err
	}

	// prompt coverage command
	coverageCmd, err := promptOptionalText("a coverage command", 200)
	if err != nil {
		return nil, err
	}

	// prompt test framework
	testFramework, err := promptOptionalText("a testing framework", 50)
	if err != nil {
		return nil, err
	}

	// prompt notes
	notes, err := promptOptionalText("any additional notes on testing", 500)
	if err != nil {
		return nil, err
	}

	// create TestingInfo
	testingInfo := &models.Testing{
		TestCommand:   testCmd,
		CoverageCmd:   coverageCmd,   // optional
		TestFramework: testFramework, // optional
		Notes:         notes,         // optional
	}

	// final validation
	if err := validation.ValidateStruct(testingInfo); err != nil {
		return nil, err
	}

	return testingInfo, nil
}

// deployment info
func PromptDeploymentInfo() (*models.Deployment, error) {
	// prompt platform
	platform, err := promptFromOptions("Choose a deployment platform", []string{"Docker", "Heroku", "AWS", "GCP", "Azure", "Vercel", "Netlify"})
	if err != nil {
		return nil, err
	}

	// prompt build command
	buildCmd, err := promptOptionalText("build command", 200)
	if err != nil {
		return nil, err
	}

	// prompt deploy command
	deployCmd, err := promptOptionalText("deploy command", 200)
	if err != nil {
		return nil, err
	}

	// prompt health check
	healthCheck, err := promptOptionalText("health check info", 200)
	if err != nil {
		return nil, err
	}

	// prompt notes
	notes, err := promptOptionalText("any additional notes on deployment", 500)
	if err != nil {
		return nil, err
	}

	// create DeploymentInfo
	deploymentInfo := &models.Deployment{
		Platform:    platform,
		BuildCmd:    buildCmd,    // optional
		DeployCmd:   deployCmd,   // optional
		HealthCheck: healthCheck, // optional
		Notes:       notes,       // optional
	}

	// final validation
	if err := validation.ValidateStruct(deploymentInfo); err != nil {
		return nil, err
	}

	return deploymentInfo, nil
}
