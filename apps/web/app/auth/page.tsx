'use client'

import { useState } from "react";
import SignInForm from "./_components/signIn/signIn";
import SignUpForm from "./_components/signUp/signUp";

export default function AuthPage() {
    const [mode, setMode] = useState<'signin' | 'signup'>('signin');
    const [userType, setUserType] = useState<'customer' | 'handyman'>('customer');

    return (
        <div className="min-h-screen bg-linear-to-br from-slate-50 via-blue-50 to-slate-100 flex items-center justify-center p-4">
            <div className="w-full max-w-md">

                <div className="text-center mb-8">
                    <h1 className="text-3xl font-bold text-slate-900 tracking-tight">QuickHand</h1>
                    <p className="text-slate-500 mt-1 text-sm">Your trusted service marketplace</p>
                </div>

                <div className="bg-white rounded-2xl shadow-xl overflow-hidden">
                    <div className="h-1 bg-linear-to-r from-blue-500 to-blue-700" />

                    <div className="p-8">
                        {/* Role tabs */}
                        <div className="flex bg-slate-100 rounded-xl p-1 mb-6">
                            <button
                                onClick={() => setUserType('customer')}
                                className={`flex-1 py-2.5 text-sm font-semibold rounded-lg transition-all duration-200 ${
                                    userType === 'customer'
                                        ? 'bg-white text-blue-600 shadow-sm'
                                        : 'text-slate-500 hover:text-slate-700'
                                }`}
                            >
                                Customer
                            </button>
                            <button
                                onClick={() => setUserType('handyman')}
                                className={`flex-1 py-2.5 text-sm font-semibold rounded-lg transition-all duration-200 ${
                                    userType === 'handyman'
                                        ? 'bg-white text-blue-600 shadow-sm'
                                        : 'text-slate-500 hover:text-slate-700'
                                }`}
                            >
                                Handyman
                            </button>
                        </div>

                        <h2 className="text-xl font-bold text-slate-900 mb-6">
                            {mode === 'signin' ? 'Welcome back' : 'Create your account'}
                        </h2>

                        {mode === 'signin'
                            ? <SignInForm key={userType} userType={userType} />
                            : <SignUpForm key={userType} userType={userType} />
                        }

                        <p className="text-center text-sm text-slate-500 mt-6">
                            {mode === 'signin' ? "Don't have an account? " : "Already have an account? "}
                            <button
                                onClick={() => setMode(mode === 'signin' ? 'signup' : 'signin')}
                                className="text-blue-600 font-semibold hover:underline"
                            >
                                {mode === 'signin' ? 'Sign up' : 'Sign in'}
                            </button>
                        </p>
                    </div>
                </div>
            </div>
        </div>
    );
}
