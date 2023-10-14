package com.vtb.hack.hack.utils;


public class CalculateDistance {

    public class Pos {
        double longitude;
        double latitude;
    }

    public static double calculateDistance(Pos myPos, Pos pointPos) {
        return Math.sqrt(Math.pow(myPos.latitude - pointPos.latitude, 2) + Math.pow(myPos.longitude - pointPos.longitude, 2));
    }
}
