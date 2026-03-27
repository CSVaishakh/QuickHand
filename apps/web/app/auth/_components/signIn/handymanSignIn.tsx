'use client'

import { ChangeEvent, FormEvent, useState } from "react";
import { useRouter } from "next/navigation";

function HandymanSignInForm() {
    const router = useRouter();
    const [formData, setFormData] = useState({ email: '', password: '' });
    const [loading, setLoading] = useState(false);
    const [error, setError] = useState('');
    const [success, setSuccess] = useState(false);

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        const { name, value } = e.target;
        setFormData((prev) => ({ ...prev, [name]: value }));
    };

    const handleSubmit = async (e: FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        setError('');
        setLoading(true);

        try {
            const response = await fetch('http://localhost:3000/handyman/sign-in', {
                method: 'POST',
                headers: { 'Content-Type': 'application/json' },
                body: JSON.stringify(formData)
            });

            if (!response.ok) {
                const contentType = response.headers.get('content-type');
                let errorMessage = 'Sign in failed';

                if (contentType?.includes('application/json')) {
                    const errorData = await response.json();
                    errorMessage = errorData.message || errorMessage;
                } else {
                    errorMessage = await response.text() || errorMessage;
                }
                throw new Error(errorMessage);
            }

            setSuccess(true);
            setFormData({ email: '', password: '' });

            setTimeout(() => {
                router.push('/');
            }, 1500);
        } catch (err) {
            setError(err instanceof Error ? err.message : 'An error occurred');
        } finally {
            setLoading(false);
        }
    };

    return (
        <section>
            {error && (
                <div className="fixed top-0 left-1/2 transform -translate-x-1/2 z-50 mt-4 animate-in slide-in-from-top duration-300">
                    <div className="bg-red-50 border-2 border-red-500 text-red-700 px-6 py-4 rounded-lg shadow-xl flex items-center gap-3">
                        <svg className="w-5 h-5 shrink-0" fill="currentColor" viewBox="0 0 20 20">
                            <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z" clipRule="evenodd" />
                        </svg>
                        <span className="font-semibold">{error}</span>
                    </div>
                </div>
            )}

            {success && (
                <div className="fixed top-0 left-1/2 transform -translate-x-1/2 z-50 mt-4 animate-in slide-in-from-top duration-300">
                    <div className="bg-green-50 border-2 border-green-500 text-green-700 px-6 py-4 rounded-lg shadow-xl flex items-center gap-3">
                        <svg className="w-5 h-5 shrink-0" fill="currentColor" viewBox="0 0 20 20">
                            <path fillRule="evenodd" d="M10 18a8 8 0 100-16 8 8 0 000 16zm3.707-9.293a1 1 0 00-1.414-1.414L9 10.586 7.707 9.293a1 1 0 00-1.414 1.414l2 2a1 1 0 001.414 0l4-4z" clipRule="evenodd" />
                        </svg>
                        <span className="font-semibold">Sign in successful! Redirecting...</span>
                    </div>
                </div>
            )}

            <div className="bg-white rounded-xl shadow-lg p-10 max-w-sm w-full">
                <h1 className="text-3xl font-bold text-blue-600 mb-8 text-center">Handyman Sign In</h1>
                <form onSubmit={handleSubmit} className="space-y-5">
                    <div>
                        <label htmlFor="email" className="block text-blue-600 text-sm font-semibold uppercase tracking-wide mb-2">Email</label>
                        <input
                            type="email"
                            id="email"
                            name="email"
                            required
                            placeholder="johndoe@gmail.com"
                            value={formData.email}
                            onChange={handleChange}
                            className="w-full px-4 py-3 border-2 border-blue-100 rounded-lg focus:outline-none focus:border-blue-600 focus:bg-white focus:ring-4 focus:ring-blue-100 transition-all bg-blue-50 text-gray-900"
                        />
                    </div>

                    <div>
                        <label htmlFor="password" className="block text-blue-600 text-sm font-semibold uppercase tracking-wide mb-2">Password</label>
                        <input
                            type="password"
                            id="password"
                            name="password"
                            required
                            placeholder="securePassword@123"
                            value={formData.password}
                            onChange={handleChange}
                            className="w-full px-4 py-3 border-2 border-blue-100 rounded-lg focus:outline-none focus:border-blue-600 focus:bg-white focus:ring-4 focus:ring-blue-100 transition-all bg-blue-50 text-gray-900"
                        />
                    </div>

                    <button
                        type="submit"
                        disabled={loading}
                        className="w-full py-3 mt-3 bg-linear-to-r from-blue-600 to-blue-700 text-white font-semibold uppercase tracking-wide rounded-lg hover:shadow-lg hover:-translate-y-0.5 active:translate-y-0 disabled:opacity-60 disabled:cursor-not-allowed transition-all duration-300"
                    >
                        {loading ? 'Signing in...' : 'Sign In'}
                    </button>
                </form>
            </div>
        </section>
    );
}

export default HandymanSignInForm;
