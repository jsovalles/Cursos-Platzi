package com.platzi.jobsearch.api;

import com.platzi.jobsearch.JobPosition;
import feign.Headers;
import feign.QueryMap;
import feign.RequestLine;

import java.util.List;
import java.util.Map;

@Headers("Accept: application/json")
public interface IApiJobs {

    @RequestLine("GET /positions.json")
    List<JobPosition> jobsAvailable(@QueryMap Map<String, Object> querymap);
}
