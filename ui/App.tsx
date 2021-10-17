import { StatusBar } from "expo-status-bar";
import React from "react";
import { StyleSheet, Text, View } from "react-native";
import { NativeRouter, Switch, Route } from "react-router-native";

import HomeComponent from "./components/Home/Home";
import ProfileComponent from "./components/Profile/Profile";
import SettingsComponent from "./components/Settings/Settings";
import FooterComponent from "./components/Layout/Footer";
import HeaderComponent from "./components/Layout/Header";
import WorkoutsComponent from "./components/Workouts/Workouts";
import TimerComponent from "./components/Timer/Timer";
import CreateCustomWorkoutComponent from "./components/Workouts/CustomWorkouts/CreateCustomWorkout";
import ProgressCalendarComponent from "./components/Progress/ProgressCalendar";

const styles = StyleSheet.create({
  container: {
    margin: 0,
    backgroundColor: "#aaa",
    height: "100%",
  }
});

export default function App() {
  return (
    <NativeRouter>
      <View style={styles.container}>
        <Switch>
          <Route component={HeaderComponent} />
        </Switch>
        <Switch>
          <Route exact path="/" component={HomeComponent} />
          <Route exact path="/profile" component={ProfileComponent} />
          <Route exact path="/settings" compoennt={SettingsComponent} />
          <Route exact path="/workouts" component={WorkoutsComponent} />
          <Route exact path="/workouts/create-workout" component={CreateCustomWorkoutComponent} />
          <Route exact path="/timer" component={TimerComponent} />
          <Route exact path="/progress" component={ProgressCalendarComponent} />
        </Switch>
        <Switch>
          <Route component={FooterComponent} />
        </Switch>
      </View>
    </NativeRouter>
  );
};
