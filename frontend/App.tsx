import { StatusBar } from "expo-status-bar";
import React from "react";
import { View, StyleSheet } from "react-native";
import Inventory from "./src/Inventory";
import UserScreen from "./src/UserScreen";
import { ApolloClient, InMemoryCache, ApolloProvider } from "@apollo/client";
import { NavigationContainer } from "@react-navigation/native";
import { createNativeStackNavigator } from "@react-navigation/native-stack";

const client = new ApolloClient({
  // TODO Set up CI/CD and map this to an internet IP
  uri: "http://192.168.0.26:8181/query",
  cache: new InMemoryCache(),
});

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <ApolloProvider client={client}>
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen
            name="Inventory"
            component={Inventory}
            options={{ title: "Inventory" }}
          />
          <Stack.Screen
            name="Users"
            component={UserScreen}
            options={{ title: "Users" }}
          />
        </Stack.Navigator>
        <View style={styles.container}>
          <UserScreen />
          <StatusBar style="auto" />
        </View>
      </NavigationContainer>
    </ApolloProvider>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: "#fff",
    alignItems: "center",
    justifyContent: "center",
  },
});
