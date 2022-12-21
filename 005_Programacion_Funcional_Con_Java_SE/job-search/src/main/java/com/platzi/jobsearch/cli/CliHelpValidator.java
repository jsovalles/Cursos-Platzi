package com.platzi.jobsearch.cli;

import com.beust.jcommander.IParameterValidator;
import com.beust.jcommander.ParameterException;

public class CliHelpValidator implements IParameterValidator {
    @Override
    public void validate(String name, String value) throws ParameterException {
        boolean actualValue = Boolean.valueOf(value);
        if(actualValue){
            throw new ParameterException("Ayuda solicitada");
        }
    }
}
