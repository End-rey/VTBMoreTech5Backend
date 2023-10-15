package com.vtb.hack.hack.utils;

import org.pytorch.IValue;
import org.pytorch.Module;
import org.pytorch.Tensor;

public class Predict {

    public static float optimize(double arg1, double arg2, double arg3) {
        Module model = Module.load("./python/optimization_algorithm/model.onnx");
        
        double[] inputData = {arg1, arg2, arg3};

        Tensor inputTensor = Tensor.fromBlob(inputData, new long[] { 1, 3 }); // Adjust shape as needed

        // Make a prediction
        IValue input = IValue.from(inputTensor);
        IValue output = model.forward(input);

        // Get the result
        Tensor outputTensor = output.toTensor();
        float[] resultData = outputTensor.getDataAsFloatArray();

        return resultData[0];
    }

}
