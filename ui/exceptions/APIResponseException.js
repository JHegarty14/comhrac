function APIResponseException(response, error) {
    this.status = response.status;
    this.error = error.errors || error;
};

export default APIResponseException;
