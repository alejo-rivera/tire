import { StatusBar } from 'expo-status-bar';
import React from 'react';
import { View, StyleSheet } from "react-native";
import Inventory from "./src/inventory";
import {
  ApolloClient,
  InMemoryCache,
  ApolloProvider,
} from "@apollo/client";
import { NavigationContainer } from '@react-navigation/native';
import { createNativeStackNavigator } from '@react-navigation/native-stack';

const client = new ApolloClient({
  // TODO Set up CI/CD and map this to an internet IP 
  uri: 'http://192.168.40.17:8181/query',
  cache: new InMemoryCache()
});

const Stack = createNativeStackNavigator();

export default function App() {
  return (
    <ApolloProvider client={client}>
      <NavigationContainer>
        <Stack.Navigator>
          <Stack.Screen name="Home" component={HomeScreen} options={{ title: 'Welcome' }} />
          <Stack.Screen name="Profile" component={ProfileScreen} />
        </Stack.Navigator>
        <View style={styles.container}>
          <Inventory />
          <StatusBar style="auto" />
        </View>
      </NavigationContainer>
    </ApolloProvider>
  );
}
const styles = StyleSheet.create({
  container: {
    flex: 1,
    backgroundColor: '#fff',
    alignItems: 'center',
    justifyContent: 'center',
  },
});
